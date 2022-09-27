package usecase

import (
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/domain/vo"
	"fmt"
)

func HandleMessage(socket service.ISocket, chatroom service.IChatRoom, roomID string, userID string, msg string) error {
	ok := socket.UserExists(roomID, userID)
	if !ok {
		return fmt.Errorf("user[%s] has not been found in room[%s]", userID, roomID)
	}
	message := vo.Message{
		From:    userID,
		To:      roomID,
		Content: msg,
	}
	chatroom.SaveMessage(roomID, message)
	socket.Broadcast(message)
	return nil
}
