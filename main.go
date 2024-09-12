package main

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/api"
	"github.com/BrunoPolaski/go-chat-app/internal/app/thirdparty/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Engine{}

	mysql.Init()

	api.Init(&engine.RouterGroup)

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
