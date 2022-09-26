package user

import (
	"fmt"
	"net/http"
	"chatroom-demo/internal/app/application/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService service.IUserService
}

func NewController(service service.IUserService) *Controller {
	return &Controller{userService: service}
}

func (ctrl Controller) Create(c *gin.Context) {
	rawBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid request",
		})
	}
	fmt.Println(string(rawBody))

	args := &CreateUserArgs{}
	args.LoadFromJSON(rawBody)
	user, err := ctrl.userService.Create(args.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func (ctrl Controller) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.userService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Internal system error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func (ctrl Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.userService.Delete(id)
	if err != nil {
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
