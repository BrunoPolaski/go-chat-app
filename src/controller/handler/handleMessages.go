package handler

import "github.com/BrunoPolaski/go-chat-app/src/domain/entity"

func HandleMessages() {
	for {
		msg := <-entity.Broadcast
		for client := range entity.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				delete(entity.Clients, client)
				client.Close()
			}
		}
	}
}
