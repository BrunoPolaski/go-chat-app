package entity

import "sync"

type Server struct {
	clients map[string]*Client
	mutex   sync.Mutex
}

var ServerInstance = &Server{
	clients: make(map[string]*Client),
}
