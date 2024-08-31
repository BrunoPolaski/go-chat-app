package entity

type User interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
}

type userDomain struct {
	id       string
	name     string
	email    string
	password string
}

func (u *userDomain) GetID() string {
	return u.id
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func NewUser(id, name, email, password string) User {
	return &userDomain{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}
