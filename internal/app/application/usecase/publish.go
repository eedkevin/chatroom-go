package usecase

import "chatroom-demo/internal/app/application"

func Publish(room application.Room, msg application.Message) {
	for userID, conn := range room.WSConns {
		if string(userID) == msg.To {
			conn.WriteJSON(msg.Content)
		}
	}
}
