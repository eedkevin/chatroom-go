package service

import (
	"net/http"
)

type ISocket interface {
	HandleConnection(w http.ResponseWriter, r *http.Request, roomID string, userID string)
	HandleMessage(roomID string, userID string, msg string) error
}
