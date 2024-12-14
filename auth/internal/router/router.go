package router

import (
	"auth/internal/handler"
	"github.com/gin-gonic/gin"
)

// SetupRoutes ...
func SetupRoutes(app *gin.Engine, handlers *handler.Handlers) *gin.Engine {
	app.POST(
		"/login",
		handlers.LoginHandler.Login,
	)

	return app
}
