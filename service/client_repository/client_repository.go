package client_repository

import (
	"github.com/gofrs/uuid"
	"model/entity"
	"service/repository"
)

func Create(client entity.ClientEntity) entity.ClientEntity {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	client.Id = id

	_, err = repository.DB.Exec(
		"INSERT INTO clients (id, code, secret) VALUES ($1, $2, $3)",
		client.Id, client.Code, client.Secret,
	)
	if err != nil {
		panic(err)
	}

	return client
}

func GetById(id uuid.UUID) *entity.ClientEntity {
	rows, err := repository.DB.Query(
		"SELECT id, code, secret FROM clients WHERE id = $1",
		id,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		client := new(entity.ClientEntity)
		err = rows.Scan(&client.Id, &client.Code, &client.Secret)
		if err != nil {
			panic(err)
		}
		return client
	}
	return nil
}
