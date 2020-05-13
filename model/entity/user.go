package entity

import "github.com/gofrs/uuid"

type UserEntity struct {
	Id       uuid.UUID
	Login    string
	Password string
	Salt     string
}
