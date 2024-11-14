package handler

import (
	"fmt"

	entity "github.com/BrunoPolaski/go-chat-app/internal/domain/entity/http"
	"github.com/BrunoPolaski/go-chat-app/internal/infra/thirdparty/thirdparty_contract"
	"github.com/gorilla/websocket"
)

type HandleMessages struct {
	logger thirdparty_contract.LoggerContract
}

func NewHandleMessages(logger thirdparty_contract.LoggerContract) HandleMessages {
	return HandleMessages{logger: logger}
}

func (hm *HandleMessages) handle(senderID, message string) {
	entity.ServerInstance.Mutex.Lock()
	defer entity.ServerInstance.Mutex.Unlock()

	for id, client := range entity.ServerInstance.Clients {
		if id != senderID {
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", senderID, message)))
			if err != nil {
				hm.logger.Error(fmt.Sprintf("Could not send message to %s, error: %v", id, err))
			}
		}
	}
}
