package service

import (
	"chatroom-demo/internal/app/domain/vo"
	"net/http"

	"github.com/gorilla/websocket"
)

type ISocket interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	Broadcast(message vo.Message)
	Register(roomID, userID string, conn *websocket.Conn)
	Deregister(roomID, userID string)
	UserExists(roomID string, userID string) bool
	Loop()
}
