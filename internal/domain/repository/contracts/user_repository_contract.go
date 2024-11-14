package contracts

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
)

type UserRepositoryContract interface {
	FindUser(email string) (entity.User, error)
	CreateUser(user entity.User) *utility.RestErr
}
