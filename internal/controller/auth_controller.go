package auth

import (
	"github.com/BrunoPolaski/go-chat-app/internal/controller/request"
	"github.com/BrunoPolaski/go-chat-app/internal/service"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (lc *AuthController) SignIn(c *gin.Context) {
	var user request.LoginRequest
	var ok bool
	if user.Username, user.Password, ok = c.Request.BasicAuth(); !ok {
		restErr := utility.NewBadRequestError("invalid auth")
		c.JSON(restErr.Code, restErr)
		return
	}

	token, err := lc.authService.SignIn(user.Username, user.Password)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, gin.H{"accessToken": token})
}
