package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type MultiRoomConns map[string]map[string]*websocket.Conn // roomID -> userID -> *websocket.Conn

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type Websocket struct {
	Upgrader         websocket.Upgrader
	BroadcastChannel chan Message
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

func (m MultiRoomConns) GetUserConn(roomID, userID string) *websocket.Conn {
	if _, ok := m[roomID]; ok {
		return m[roomID][userID]
	} else {
		return nil
	}
}

func (m MultiRoomConns) GetUserConnsInRoom(roomID string) map[string]*websocket.Conn {
	if _, ok := m[roomID]; ok {
		return m[roomID]
	} else {
		return nil
	}
}

func (s *Websocket) HandleConnection(w http.ResponseWriter, r *http.Request, roomID string, userID string) {
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error on upgrading http to ws, room[%s] user[%s], err: %v", roomID, userID, err)
	}
	defer conn.Close()

	s.Clients.Register(roomID, userID, conn)
	log.Printf("client[%s] connected to room[%s]", userID, roomID)
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(s.Clients[roomID], userID)
			log.Printf("error on reading message from client, %v", err)
			log.Printf("client[%s] disconnected from room[%s]", userID, roomID)
			break
		}
		log.Printf("received message from client[%s]: %s", msg.From, msg.Content)
		s.BroadcastChannel <- msg
	}
}

func (s *Websocket) HandleMessage(roomID string, userID string, msg string) error {
	ok := s.Clients.UserExists(roomID, userID)
	if !ok {
		return fmt.Errorf("user[%s] has not been found in room[%s]", userID, roomID)
	}
	message := Message{
		From:    userID,
		To:      roomID,
		Content: msg,
	}

	s.BroadcastChannel <- message
	return nil
}

func (s *Websocket) Loop() {
	for {
		msg := <-s.BroadcastChannel
		roomID := msg.To
		userConns := s.Clients.GetUserConnsInRoom(roomID)
		for userID, conn := range userConns {
			err := conn.WriteJSON(msg)
			if err != nil {
				conn.Close()
				s.Clients.Deregister(roomID, userID)
			}
		}
	}
}
