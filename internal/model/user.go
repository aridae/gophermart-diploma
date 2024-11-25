package model

import "golang.org/x/crypto/bcrypt"

type UserCredentials struct {
	Login        string
	PasswordHash []byte
}

func NewUserCredentials(login string, pass string) (UserCredentials, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return UserCredentials{}, err
	}

	return UserCredentials{
		Login:        login,
		PasswordHash: pwdHash,
	}, nil
}

func (c UserCredentials) Equal(login string, pass string) bool {
	if c.Login != login {
		return false
	}

	if err := bcrypt.CompareHashAndPassword(c.PasswordHash, []byte(pass)); err != nil {
		return false
	}

	return true
}

type User struct {
	Login string
}
