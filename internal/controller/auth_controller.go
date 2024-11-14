package controller

import (
	"net/http"

	"github.com/BrunoPolaski/go-chat-app/internal/controller/request"
	"github.com/BrunoPolaski/go-chat-app/internal/domain/service"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) authController {
	return authController{
		authService: authService,
	}
}

func (lc *authController) SignIn(c *gin.Context) {
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"accessToken": token,
		},
	)
}
