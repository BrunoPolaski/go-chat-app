package service

import "github.com/BrunoPolaski/go-chat-app/pkg/utility"

type AuthService interface {
	SignUp(name, email, password string) *utility.RestErr
	SignIn(email, password string) (string, *utility.RestErr)
}

type authService struct {
	authRepository AuthRepository
}
