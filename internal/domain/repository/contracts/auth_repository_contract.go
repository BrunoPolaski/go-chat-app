package contracts

type AuthRepositoryContract interface {
	FindUser(username, password string) (bool, error)
}
