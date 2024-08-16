package routes

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/src/handler"
)

func Init() {
	http.HandleFunc("/ws", handler.HandleConnections)
}
