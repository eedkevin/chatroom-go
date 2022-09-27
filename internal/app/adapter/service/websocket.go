package service

import (
	"chatroom-demo/internal/app/domain/vo"
	"net/http"

	"github.com/gorilla/websocket"
)

type MultiRoomConns map[string]map[string]*websocket.Conn // roomID -> userID -> *websocket.Conn

type Message struct {
	RoomID  string `json:"room_id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type Websocket struct {
	Upgrader         websocket.Upgrader
	BroadcastChannel chan Message
	PublishChannel   chan Message
	Clients          MultiRoomConns
}

func NewWebsocket() *Websocket {
	return &Websocket{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
		Clients:          make(MultiRoomConns),
		BroadcastChannel: make(chan Message),
		PublishChannel:   make(chan Message),
	}
}

func (m MultiRoomConns) Register(roomID, userID string, conn *websocket.Conn) {
	if _, ok := m[roomID]; ok {
		m[roomID][userID] = conn
	} else {
		m[roomID] = make(map[string]*websocket.Conn)
		m[roomID][userID] = conn
	}
}

func (m MultiRoomConns) Deregister(roomID, userID string) {
	if _, ok := m[roomID]; ok {
		delete(m[roomID], userID)
	}
}

func (m MultiRoomConns) UserExists(roomID, userID string) (ok bool) {
	_, ok = m[roomID]
	return
}

func (m MultiRoomConns) GetUserConn(roomID, userID string) (*websocket.Conn, bool) {
	if _, ok := m[roomID]; ok {
		if _, ok := m[roomID][userID]; ok {
			return m[roomID][userID], true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

func (m MultiRoomConns) GetUserConnsInRoom(roomID string) (map[string]*websocket.Conn, bool) {
	if _, ok := m[roomID]; ok {
		return m[roomID], true
	} else {
		return nil, false
	}
}

func (s *Websocket) Broadcast(roomID string, message vo.Message) {
	msg := Message{
		RoomID:  roomID,
		From:    message.From,
		To:      message.To,
		Content: message.Content,
	}
	s.BroadcastChannel <- msg
}

func (s *Websocket) Publish(roomID string, userID string, message vo.Message) {
	msg := Message{
		RoomID:  roomID,
		From:    message.From,
		To:      userID,
		Content: message.Content,
	}
	s.PublishChannel <- msg
}

func (s *Websocket) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return s.Upgrader.Upgrade(w, r, nil)
}

func (s *Websocket) Register(roomID, userID string, conn *websocket.Conn) {
	s.Clients.Register(roomID, userID, conn)
}

func (s *Websocket) Deregister(roomID, userID string) {
	s.Clients.Deregister(roomID, userID)
}

func (s *Websocket) UserExists(roomID string, userID string) bool {
	return s.Clients.UserExists(roomID, userID)
}

func (s *Websocket) Loop() {
	for {
		select {
		case msg := <-s.BroadcastChannel:
			roomID := msg.RoomID
			userConns, ok := s.Clients.GetUserConnsInRoom(roomID)
			if ok {
				for userID, conn := range userConns {
					err := conn.WriteJSON(msg)
					if err != nil {
						conn.Close()
						s.Clients.Deregister(roomID, userID)
					}
				}
			}
		case msg := <-s.PublishChannel:
			roomID := msg.RoomID
			userID := msg.To
			conn, ok := s.Clients.GetUserConn(roomID, userID)
			if ok {
				err := conn.WriteJSON(msg)
				if err != nil {
					conn.Close()
					s.Clients.Deregister(roomID, userID)
				}
			}
		}
	}
}
