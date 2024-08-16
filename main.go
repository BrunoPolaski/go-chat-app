package main

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/src/handler"
	"github.com/BrunoPolaski/go-chat-app/src/routes"
)

func main() {
	routes.Init()
	go handler.HandleMessages()

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
