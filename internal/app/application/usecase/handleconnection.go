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

	historicalMsg, err := chatroom.ListMessages(roomID)
	if err != nil {
		log.Printf("error on getting historical messages for user[%s] from room[%s]", userID, roomID)
	}
	go func() {
		log.Printf("sync historical messages to newly joined user[%s] in room[%s]", userID, roomID)
		for _, msg := range historicalMsg {
			socket.Publish(roomID, userID, msg)
		}
	}()

	for {
		var msg vo.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			socket.Deregister(roomID, userID)
			log.Printf("error on reading message from client[%s], disconnected, %v", userID, err)
			break
		}
		log.Printf("received message from client[%s]: %s", msg.From, msg.Content)
		go chatroom.SaveMessage(roomID, msg)
		go socket.Broadcast(roomID, msg)
	}
}
