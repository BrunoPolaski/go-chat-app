package repository

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity"
	"github.com/BrunoPolaski/go-chat-app/internal/domain/entity/dto"
	"github.com/BrunoPolaski/go-chat-app/internal/infra/thirdparty/thirdparty_contract"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
)

type authRepository struct {
	database thirdparty_contract.DatabaseContract
}

func NewAuthRepository(database thirdparty_contract.DatabaseContract) *authRepository {
	return &authRepository{
		database: database,
	}
}

func (ar *authRepository) FindUser(email string) (entity.User, error) {
	conn, err := ar.database.Connect()
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
