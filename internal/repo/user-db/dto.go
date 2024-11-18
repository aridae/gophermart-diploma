package userdb

import "github.com/aridae/gophermart-diploma/internal/model"

type userDTO struct {
	ID           int64  `db:"id"`
	Login        string `db:"login"`
	PasswordHash []byte `db:"pwd_hash"`
}

func mapDTOToDomainUserCredentials(dto userDTO) model.UserCredentials {
	return model.UserCredentials{
		Login:        dto.Login,
		PasswordHash: dto.PasswordHash,
	}
}
