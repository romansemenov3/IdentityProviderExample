package api

import (
	"api/controller"
	"api/index_controller"
	"api/request_handler"
	"api/user_controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Routes []controller.Route

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handle(route))
	}

	return router
}

func handle(route controller.Route) http.Handler {
	var handler http.Handler
	handler = route.HandlerFunc

	handler = request_handler.DisableAccessControl(handler)
	handler = request_handler.TokenValidator(handler, route.Grants)
	handler = request_handler.AccessLogger(handler, route.Name)
	handler = request_handler.ErrorAdvice(handler)

	return handler
}

var routes = Routes{
	index_controller.RouteGet,

	user_controller.RouteCreate,
	user_controller.RouteGetById,
	user_controller.RouteDeleteById,
	user_controller.RouteLogin,
}
