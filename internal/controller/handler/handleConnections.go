package handler

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/internal/thirdparty/contract"
)

type HandleConnections struct {
	logger         contract.LoggerContract
	handleMessages HandleMessages
}

func NewHandleConnections(logger contract.LoggerContract, handleMessages HandleMessages) HandleConnections {
	return HandleConnections{
		logger:         logger,
		handleMessages: handleMessages,
	}
}

func (hc *HandleConnections) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := entity.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		hc.logger.Error(fmt.Sprintf("Could not upgrade connection, error: %v", err))
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
			hc.logger.Error(fmt.Sprintf("Could not read message, error: %v", err))
			break
		}

		hc.handleMessages.handle(client.ID, string(message))
	}
}
