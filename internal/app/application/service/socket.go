package service

import (
	"chatroom-demo/internal/app/domain/vo"
	"net/http"

	"github.com/gorilla/websocket"
)

type ISocket interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	Broadcast(roomID string, message vo.Message)
	Publish(roomID string, userID string, message vo.Message)
	Register(roomID, userID string, conn *websocket.Conn)
	Deregister(roomID, userID string)
	RoomExists(roomID string) bool
	UserExists(roomID string, userID string) bool
	Loop()
}
