package identity_provider_error

import "net/http"

type ClientNotFoundError struct {
	Id string
}

func (e ClientNotFoundError) Error() string {
	return "Client with id " + e.Id + " not found"
}

func (e ClientNotFoundError) Title() string {
	return "Entity not found"
}

func (e ClientNotFoundError) Code() string {
	return prefix + "1001"
}

func (e ClientNotFoundError) Status() int {
	return http.StatusNotFound
}
