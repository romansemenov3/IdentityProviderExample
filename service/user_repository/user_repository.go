package user_repository

import (
	"github.com/gofrs/uuid"
	"model/entity"
	"service/repository"
)

func Create(user entity.UserEntity) entity.UserEntity {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	user.Id = id

	_, err = repository.DB.Exec(
		"INSERT INTO users (id, login, password, salt) VALUES ($1, $2, $3, $4)",
		user.Id, user.Login, user.Password, user.Salt,
	)
	if err != nil {
		panic(err)
	}

	return user
}

func Update(user entity.UserEntity) {
	_, err := repository.DB.Exec(
		"UPDATE users SET password = $2, salt = $3 WHERE id = $1",
		user.Id, user.Password, user.Salt,
	)
	if err != nil {
		panic(err)
	}
}

func DeleteById(id uuid.UUID) {
	_, err := repository.DB.Exec(
		"DELETE FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		panic(err)
	}
}

func FindById(id uuid.UUID) *entity.UserEntity {
	rows, err := repository.DB.Query(
		"SELECT id, login, password, salt FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		user := new(entity.UserEntity)
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Salt)
		if err != nil {
			panic(err)
		}
		return user
	}
	return nil
}

func FindByLogin(login string) *entity.UserEntity {
	rows, err := repository.DB.Query(
		"SELECT id, login, password, salt FROM users WHERE login = $1",
		login,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		user := new(entity.UserEntity)
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Salt)
		if err != nil {
			panic(err)
		}
		return user
	}
	return nil
}
