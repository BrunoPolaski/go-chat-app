package handler

import (
	"fmt"

	"github.com/BrunoPolaski/go-chat-app/internal/app/thirdparty/logger"
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/gorilla/websocket"
)

func HandleMessages(senderID, message string) {
	entity.ServerInstance.Mutex.Lock()
	defer entity.ServerInstance.Mutex.Unlock()

	for id, client := range entity.ServerInstance.Clients {
		if id != senderID {
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", senderID, message)))
			if err != nil {
				logger.Error(fmt.Sprintf("Could not send message to %s", id), err)
			}
		}
	}
}
