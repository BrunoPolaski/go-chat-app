package repository

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/internal/dto"
	"github.com/BrunoPolaski/go-chat-app/internal/thirdparty/contract"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
)

type AuthRepository struct {
	database contract.DatabaseContract
}

func NewAuthRepository(database contract.DatabaseContract) *AuthRepository {
	return &AuthRepository{database: database}
}

func (lr *AuthRepository) FindUser(email string) (entity.User, error) {
	conn, err := lr.database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	if result.Next() {
		var userDto dto.UserDTO
		err = result.Scan(&userDto)
		if err != nil {
			return nil, err
		}

		return userDto.ToDomain(), nil
	}

	return nil, utility.NewNotFoundError("User not found")
}
