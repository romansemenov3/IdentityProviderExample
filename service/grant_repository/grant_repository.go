package grant_repository

import (
	"github.com/gofrs/uuid"
	"model/entity"
	"service/repository"
)

func Create(grant entity.GrantEntity) entity.GrantEntity {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	grant.Id = id

	_, err = repository.DB.Exec(
		"INSERT INTO grants (id, code) VALUES ($1, $2)",
		grant.Id, grant.Code,
	)
	if err != nil {
		panic(err)
	}

	return grant
}

func FindByUserId(userId uuid.UUID) []entity.GrantEntity {
	rows, err := repository.DB.Query(
		"SELECT id, code FROM grants WHERE id IN (SELECT grant_id FROM users_grants WHERE user_id = $1)",
		userId,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	grants := make([]entity.GrantEntity, 0)
	for rows.Next() {
		grant := entity.GrantEntity{}
		err := rows.Scan(&grant.Id, &grant.Code)
		if err != nil {
			panic(err)
		}
		grants = append(grants, grant)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return grants
}
