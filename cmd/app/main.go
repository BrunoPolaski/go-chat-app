package main

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/api"
	"github.com/BrunoPolaski/go-chat-app/internal/infra/thirdparty/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.NewZapLoggerAdapter()

	r := gin.Default()
	api.Init(&r.RouterGroup, logger)

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
