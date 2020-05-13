package entity

import "github.com/gofrs/uuid"

type GrantEntity struct {
	Id   uuid.UUID
	Code string
}
