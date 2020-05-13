package controller

import "net/http"

const Prefix = "/api/identity-provider/v1"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Grants      []string
}
