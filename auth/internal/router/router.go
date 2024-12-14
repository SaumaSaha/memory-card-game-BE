package router

import (
	"auth/internal/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, handlers *handler.Handlers) *gin.Engine {
	fmt.Println("setting up router")
	app.POST(
		"/login",
		handlers.LoginHandler.Login,
	)
	return app
}
