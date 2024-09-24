package dto

import "github.com/BrunoPolaski/go-chat-app/internal/domain/entity"

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (u *UserDTO) ToDomain() entity.User {
	return entity.NewUser(
		u.ID,
		u.Name,
		u.Email,
		u.Password,
	)
}

func (u *UserDTO) FromDomain(entity entity.User) {
	u.ID = entity.GetID()
	u.Name = entity.GetName()
	u.Email = entity.GetEmail()
	u.Password = entity.GetPassword()
}
