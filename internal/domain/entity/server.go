package entity

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
}

var ServerInstance = &Server{
	Clients: make(map[string]*Client),
	Mutex:   sync.Mutex{},
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
