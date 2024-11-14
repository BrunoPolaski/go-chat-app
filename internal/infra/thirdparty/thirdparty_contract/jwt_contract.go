package thirdparty_contract

import "github.com/BrunoPolaski/go-chat-app/pkg/utility"

type JWTContract interface {
	GenerateToken(claims map[string]interface{}) (string, *utility.RestErr)
	ParseToken(token string) (map[string]interface{}, *utility.RestErr)
}
