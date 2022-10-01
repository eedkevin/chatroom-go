package usecase

import (
	"chatroom-demo/internal/app/application"
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/domain/vo"
	"fmt"
)

func HandleHTTPMessage(socket service.ISocket, chatroom service.IChatRoom, roomID string, userID string, msg string) error {
	ok := socket.RoomExists(roomID)
	if !ok {
		return fmt.Errorf(application.NotFoundErr)
	}

	message := vo.Message{
		From:    userID,
		To:      vo.MESSAGE_TO_ALL,
		Content: msg,
	}
	chatroom.SaveMessage(roomID, message)
	socket.Broadcast(roomID, message)
	return nil
}
