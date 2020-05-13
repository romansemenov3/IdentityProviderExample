package user_controller

import (
	"api/controller"
	"encoding/json"
	"model/dto"
	"net/http"
	"service/user_service"
)

const prefixes = controller.Prefix + "/users"

var RouteLogin = controller.Route{
	Name:        "Login",
	Method:      "POST",
	Pattern:     prefixes + "/login",
	HandlerFunc: login,
}
func login(w http.ResponseWriter, r *http.Request) {
	var login dto.FormLoginDTO
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		panic(err)
	}

	token := user_service.Login(login)
	body, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}