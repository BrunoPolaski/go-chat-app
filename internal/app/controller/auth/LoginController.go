package auth

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
