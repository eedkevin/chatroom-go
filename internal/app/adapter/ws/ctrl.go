package ws

import (
	"net/http"
	"chatroom-demo/internal/app/application/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Controller struct {
	websocketService service.ISocket
}

func NewController(service service.ISocket) *Controller {
	return &Controller{websocketService: service}
}

func (ctrl Controller) Connect(c *gin.Context) {
	ctrl.websocketService.HandleConnection(c.Writer, c.Request, c.Param("roomID"), c.Param("userID"))
}

func (ctrl Controller) Broadcast(c *gin.Context) {
	_, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}
	// ctrl.websocketService.HandleMessage(c.Writer, c.Request, c.Param("roomID"), c.Param("userID"))
}
