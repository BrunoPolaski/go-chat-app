package contracts

import "github.com/BrunoPolaski/go-chat-app/internal/domain/entity"

type UserRepositoryContract interface {
	FindUser(email string) (entity.User, error)
	CreateUser(user entity.User) error
}
