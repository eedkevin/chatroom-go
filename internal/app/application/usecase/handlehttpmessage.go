package usecase

import (
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/domain/vo"
)

func HandleHTTPMessage(socket service.ISocket, chatroom service.IChatRoom, roomID string, userID string, msg string) error {

	message := vo.Message{
		From:    userID,
		To:      vo.MESSAGE_TO_ALL,
		Content: msg,
	}
	chatroom.SaveMessage(roomID, message)
	socket.Broadcast(roomID, message)
	return nil
}
