package client_service

import (
	"encoding/hex"
	"github.com/gofrs/uuid"
	"math/rand"
	"model/entity"
	"service/client_repository"
	"time"
)

func randomHex() string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func Create(code string) entity.ClientEntity {
	client := entity.ClientEntity{Code: code, Secret: randomHex()}
	return client_repository.Create(client)
}

func GetByIdOrThrow(id string) *entity.ClientEntity {
	uuid, err := uuid.FromString(id)
	if err != nil {
		panic(err)
	}
	return client_repository.GetById(uuid)
}
