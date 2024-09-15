package api

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/internal/app/controller/auth"
	"github.com/BrunoPolaski/go-chat-app/internal/app/controller/handler"
	"github.com/BrunoPolaski/go-chat-app/internal/app/thirdparty/logger"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	logger := logger.NewLoggerAdapter()
	handleMessages := handler.NewHandleMessages(logger)
	handleConnections := handler.NewHandleConnections(
		logger,
		handleMessages,
	)
	http.HandleFunc("/ws", handleConnections.Handle)
	authentication := r.Group("/auth")
	{
		authentication.POST("/login", auth.LoginController)
	}
}
