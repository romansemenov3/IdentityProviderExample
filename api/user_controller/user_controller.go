package user_controller

import (
	"api/controller"
	"encoding/json"
	"model/dto"
	"net/http"
	"service/user_service"
	"strings"
)

const prefix = controller.Prefix + "/user"

var RouteGetById = controller.Route{
	Name:        "GetById",
	Method:      "GET",
	Pattern:     prefix + "/{id}",
	HandlerFunc: getById,
	Grants: []string{"USER_MANAGER"},
}
func getById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, prefix + "/")
	user := user_service.GetUserByIdOrThrow(id)
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

var RouteDeleteById = controller.Route{
	Name:        "DeleteById",
	Method:      "DELETE",
	Pattern:     prefix + "/{id}",
	HandlerFunc: deleteById,
	Grants: []string{"USER_MANAGER"},
}
func deleteById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, prefix + "/")
	user_service.DeleteUserById(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

var RouteCreate = controller.Route{
	Name:        "Create",
	Method:      "POST",
	Pattern:     prefix,
	HandlerFunc: create,
}
func create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	user = user_service.CreateUser(user)
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
