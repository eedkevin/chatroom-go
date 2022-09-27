package app

import (
	"chatroom-demo/internal/app/adapter/http/index"
	"chatroom-demo/internal/app/adapter/http/room"
	"chatroom-demo/internal/app/adapter/http/user"
	"chatroom-demo/internal/app/adapter/repository"
	"chatroom-demo/internal/app/adapter/service"
	"chatroom-demo/internal/app/adapter/ws"
)

func setupRoute(server *Server) {
	server.StaticFile("/", "./public/index.html")

	indexGroup := server.Group("/")
	{
		indexCtrl := index.Controller{}
		indexGroup.Any("/healthz", indexCtrl.Healthz)
	}

	apiGroup := server.Group("/api")
	{
		wsService := service.NewWebsocket()
		go wsService.Loop()

		chatService := service.NewChatRoomService(repository.NewRoomRepo(server.Storage["rooms"]))
		userService := service.NewUserService(repository.NewUserRepo(server.Storage["users"]))

		roomGroup := apiGroup.Group("/rooms")
		{
			roomCtrl := room.NewController(chatService, wsService)
			roomGroup.POST("/", roomCtrl.Create)
			roomGroup.GET("/", roomCtrl.List)
			roomGroup.GET("/:id", roomCtrl.Get)
			roomGroup.GET("/:id/thumbnail", roomCtrl.Thumbnail)
			roomGroup.DELETE("/:id", roomCtrl.Destroy)
			roomGroup.POST("/:id/publish", roomCtrl.Publish)
		}

		userGroup := apiGroup.Group("/users")
		{
			userCtrl := user.NewController(userService)
			userGroup.POST("/", userCtrl.Create)
			userGroup.GET("/:id", userCtrl.Get)
			userGroup.DELETE("/:id", userCtrl.Delete)
		}

		wsGroup := apiGroup.Group("/ws")
		{
			wsCtrl := ws.NewController(wsService, chatService)
			wsGroup.GET("/:roomID/:userID", wsCtrl.Connect)
		}

	}

}
