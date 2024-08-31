package bcrypt

type Bcrypt interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}
