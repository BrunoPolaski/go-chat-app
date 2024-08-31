package api

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/internal/app/controller/auth"
	"github.com/BrunoPolaski/go-chat-app/internal/app/controller/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	http.HandleFunc("/ws", handler.HandleConnections)
	authentication := r.Group("/auth")
	{
		authentication.POST("/login", auth.Login)
	}
}
