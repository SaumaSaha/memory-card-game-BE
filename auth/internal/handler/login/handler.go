package login

import "github.com/gin-gonic/gin"

// Handler ...
type Handler interface {
	Login(ctx *gin.Context)
}
