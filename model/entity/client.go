package entity

import "github.com/gofrs/uuid"

type ClientEntity struct {
	Id     uuid.UUID
	Code   string
	Secret string
}
