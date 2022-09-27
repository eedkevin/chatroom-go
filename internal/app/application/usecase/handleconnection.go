package usecase

import (
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/domain/vo"
	"log"
	"net/http"
)

func HandleConnection(socket service.ISocket, chatroom service.IChatRoom, w http.ResponseWriter, r *http.Request, roomID string, userID string) {
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		log.Printf("error on upgrading http to ws, room[%s] user[%s], err: %v", roomID, userID, err)
	}
	defer conn.Close()

	socket.Register(roomID, userID, conn)
	log.Printf("client[%s] connected to room[%s]", userID, roomID)
	for {
		var msg vo.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			socket.Deregister(roomID, userID)
			log.Printf("error on reading message from client, %v", err)
			log.Printf("client[%s] disconnected from room[%s]", userID, roomID)
			break
		}
		log.Printf("received message from client[%s]: %s", msg.From, msg.Content)
		chatroom.SaveMessage(roomID, msg)
		socket.Broadcast(msg)
	}
}
