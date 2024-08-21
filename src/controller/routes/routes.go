package routes

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/src/application/auth"
	"github.com/BrunoPolaski/go-chat-app/src/controller/handler"
)

func Init() {
	http.HandleFunc("/ws", handler.HandleConnections)
	http.HandleFunc("/auth/login", auth.LoginController)
}
