package contract

type JWTContract interface {
	GenerateToken(claims map[string]interface{}) (string, error)
	ParseToken(token string) (map[string]interface{}, error)
}
