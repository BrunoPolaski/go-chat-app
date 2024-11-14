package bcrypt

type BcryptAdapter struct{}

func (ba *BcryptAdapter) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return nil, nil
}

func (ba *BcryptAdapter) CompareHashAndPassword(hashedPassword, password []byte) error {
	return nil
}
