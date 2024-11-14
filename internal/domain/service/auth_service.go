package service

import (
	"github.com/BrunoPolaski/go-chat-app/internal/domain/repository/contracts"
	"github.com/BrunoPolaski/go-chat-app/internal/infra/thirdparty/thirdparty_contract"
	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	authRepository contracts.AuthRepositoryContract
	jwt            thirdparty_contract.JWTContract
}

func NewAuthService(authRepository contracts.AuthRepositoryContract, jwt thirdparty_contract.JWTContract) *authService {
	return &authService{
		authRepository: authRepository,
		jwt:            jwt,
	}
}

func (as *authService) SignIn(email, password string) (string, *utility.RestErr) {
	err := as.authRepository.SignIn(email, password)
	if err != nil {
		return "", err
	}

	var token string
	token, err = as.jwt.GenerateToken(
		jwt.MapClaims{
			"email": email,
		},
	)

	if err != nil {
		return "", err
	} else {
		return token, nil
	}

}
