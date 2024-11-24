package userrepo

import "github.com/aridae/gophermart-diploma/internal/model"

type userDTO struct {
	Login        string `db:"login"`
	PasswordHash []byte `db:"pwd_hash"`
}

func mapDTOToDomainUserCredentials(dto userDTO) model.UserCredentials {
	return model.UserCredentials{
		Login:        dto.Login,
		PasswordHash: dto.PasswordHash,
	}
}
