package usecase

import (
	"log"
	"chatroom-demo/internal/app/application"
)

func Broadcast(room application.Room, msg string) {
	for userID, conn := range room.WSConns {
		if err := conn.WriteJSON(msg); err != nil {
			log.Printf("error on sending message to User[%s]", userID)
		}
	}
}
