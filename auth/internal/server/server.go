package server

import (
	"log"

	"auth/internal/handler"
	loginhandler "auth/internal/handler/login"
	"auth/internal/router"
	loginservice "auth/internal/service/login"
	"github.com/gin-gonic/gin"
)

// Init ...
func Init() {
	app := gin.Default()

	handlers := initHandlers()
	app = router.SetupRoutes(app, handlers)

	err := app.Run(":9090")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func initHandlers() *handler.Handlers {
	loginHandler := newLoginHandler()

	return &handler.Handlers{
		LoginHandler: loginHandler,
	}
}

func newLoginHandler() loginhandler.Handler {
	loginService := loginservice.NewService()

	return loginhandler.NewHandler(loginService)
}
