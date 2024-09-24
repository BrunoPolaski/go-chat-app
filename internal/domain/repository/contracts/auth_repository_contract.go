package contracts

import "github.com/BrunoPolaski/go-chat-app/internal/domain/entity"

type AuthRepositoryContract interface {
	SignUp(user entity.User) error
}
