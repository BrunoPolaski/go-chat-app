package handler

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/src/domain/entity"
	"github.com/gorilla/websocket"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer ws.Close()
	entity.Clients[ws] = true

	for {
		var msg entity.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(entity.Clients, ws)
			break
		}
		entity.Broadcast <- msg
	}
}
