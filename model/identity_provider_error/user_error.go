package identity_provider_error

import "net/http"

type UserNotFoundError struct {
	Id string
}

func (e UserNotFoundError) Error() string {
	return "User with id " + e.Id + " not found"
}

func (e UserNotFoundError) Title() string {
	return "Entity not found"
}

func (e UserNotFoundError) Code() string {
	return prefix + "1001"
}

func (e UserNotFoundError) Status() int {
	return http.StatusNotFound
}

type UserAlreadyExistsError struct {
	Login string
}

func (e UserAlreadyExistsError) Error() string {
	return "User with login " + e.Login + " already exists"
}

func (e UserAlreadyExistsError) Title() string {
	return "Entity not found"
}

func (e UserAlreadyExistsError) Code() string {
	return prefix + "1002"
}

func (e UserAlreadyExistsError) Status() int {
	return http.StatusConflict
}
