package handler

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/internal/app/thirdparty/contract"
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
)

type HandleConnections struct {
	loggerAdapter  contract.LoggerContract
	handleMessages HandleMessages
}

func (hc *HandleConnections) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := entity.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		hc.loggerAdapter.Error("Could not upgrade connection", err)
		return
	}

	// Check the database for the user
	// If the user is not found, return an error
	// If the user is found, create a new client and add it to the server

	client := &entity.Client{
		Conn: conn,
		ID:   "1",
	}

	entity.ServerInstance.Mutex.Lock()
	entity.ServerInstance.Clients[client.ID] = client
	entity.ServerInstance.Mutex.Unlock()

	defer func() {
		entity.ServerInstance.Mutex.Lock()
		delete(entity.ServerInstance.Clients, client.ID)
		entity.ServerInstance.Mutex.Unlock()
		client.Conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			hc.loggerAdapter.Error("Could not read message", err)
			break
		}

		HandleMessages(client.ID, string(message))
	}
}
