package room

import (
	"chatroom-demo/internal/app/application"
	"chatroom-demo/internal/app/application/service"
	"chatroom-demo/internal/app/application/usecase"
	"chatroom-demo/internal/app/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	chatroomService  service.IChatRoom
	websocketService service.ISocket
}

func NewController(chatroom service.IChatRoom, websocket service.ISocket) *Controller {
	return &Controller{chatroomService: chatroom, websocketService: websocket}
}

func (ctrl Controller) Create(c *gin.Context) {
	rawBody, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid request",
		})
		return
	}

	fmt.Println(string(rawBody))

	args := &CreateRoomArgs{}
	args.LoadFromJSON(rawBody)
	room, err := ctrl.chatroomService.Create(args.Name, domain.ROOM_TYPE_PUBLIC)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   room,
	})
}

func (ctrl Controller) List(c *gin.Context) {
	rooms, err := ctrl.chatroomService.List()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   rooms,
	})
}

func (ctrl Controller) Get(c *gin.Context) {
	id := c.Param("id")
	room, err := ctrl.chatroomService.Get(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   room,
	})
}

func (ctrl Controller) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.chatroomService.Destroy(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (ctrl Controller) Broadcast(c *gin.Context) {
	roomID := c.Param("id")
	rawBody, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid request",
		})
		return
	}

	args := &BroadcastMessageArgs{}
	args.LoadFromJSON(rawBody)
	err = usecase.HandleHTTPMessage(ctrl.websocketService, ctrl.chatroomService, roomID, args.From, args.Content)
	if err != nil {
		if err.Error() == application.NotFoundErr {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"msg":    "room does not exist",
			})
			return
		}

		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
	}
}

func (ctrl Controller) Thumbnail(c *gin.Context) {
	roomID := c.Param("id")
	thumbnail, err := ctrl.chatroomService.Thumbnail(roomID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   thumbnail,
	})
}
