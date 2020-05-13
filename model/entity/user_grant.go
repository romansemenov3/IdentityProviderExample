package entity

import "github.com/gofrs/uuid"

type UserGrantEntity struct {
	UserId  uuid.UUID
	GrantId uuid.UUID
}
