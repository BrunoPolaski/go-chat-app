package repository

import "github.com/BrunoPolaski/go-chat-app/internal/app/thirdparty/contract"

type AuthRepository struct {
	database contract.Database
}

func NewLoginRepository(database contract.Database) *AuthRepository {
	return &AuthRepository{database: database}
}

func (lr *AuthRepository) FindUser(username, password string) (bool, error) {
	conn, err := lr.database.GetConn()
	if err != nil {
		return false, err
	}
	defer conn.Close()

}
