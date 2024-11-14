package api

import (
	"net/http"

	auth "github.com/BrunoPolaski/go-chat-app/internal/controller"
	"github.com/BrunoPolaski/go-chat-app/internal/controller/handler"
	"github.com/BrunoPolaski/go-chat-app/internal/infra/thirdparty/thirdparty_contract"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, logger thirdparty_contract.LoggerContract) {
	handleMessages := handler.NewHandleMessages(logger)
	handleConnections := handler.NewHandleConnections(
		logger,
		handleMessages,
	)
	http.HandleFunc("/ws", handleConnections.Handle)

	authentication := r.Group("/auth")
	{
		authentication.POST("/login", auth.AuthController.SignIn)
	}
}
