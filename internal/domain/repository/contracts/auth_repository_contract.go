package contracts

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
)

type AuthRepositoryContract interface {
	SignUp(user entity.User) *utility.RestErr
	SignIn(username, password string) *utility.RestErr
}
