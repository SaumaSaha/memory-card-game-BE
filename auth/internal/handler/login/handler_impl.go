package login

import (
	"net/http"

	"auth/internal/service/login"
	"github.com/gin-gonic/gin"
)

type loginHandler struct {
	loginService login.Service
}

// NewHandler ...
func NewHandler(loginService login.Service) Handler {
	return &loginHandler{
		loginService: loginService,
	}
}

func (lh *loginHandler) Login(ctx *gin.Context) {
	var loginData struct {
		Username string `binding:"required" json:"username"`
		Password string `binding:"required" json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})

		return
	}

	if valid, _ := lh.loginService.ValidateCredentials(
		ctx.Request.Context(),
		loginData.Username,
		loginData.Password,
	); valid {
		ctx.JSON(http.StatusOK, gin.H{"token": "dummy-token"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
