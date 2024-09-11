package auth

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utility.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	
}
