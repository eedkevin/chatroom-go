package ws

import (
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Controller struct {
	websocketService service.ISocket
	chatroomService  service.IChatRoom
}

func NewController(service service.ISocket, chatroom service.IChatRoom) *Controller {
	return &Controller{websocketService: service, chatroomService: chatroom}
}

func (ctrl Controller) Connect(c *gin.Context) {
	usecase.HandleConnection(ctrl.websocketService, ctrl.chatroomService, c.Writer, c.Request, c.Param("roomID"), c.Param("userID"))
}
