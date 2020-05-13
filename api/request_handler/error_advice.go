package request_handler

import (
	"encoding/json"
	"log"
	"model/dto"
	"model/identity_provider_error"
	"net/http"
)

func ErrorAdvice(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverIfNecessary(w, r)
		inner.ServeHTTP(w, r)
	})
}

func recoverIfNecessary(w http.ResponseWriter, r *http.Request) {
	if rec := recover(); rec != nil {
		handleError(rec, w, r)
	}
}

func handleError(rec interface{}, w http.ResponseWriter, r *http.Request) {
	err, isError := rec.(error)
	if !isError {
		log.Printf("Recovered %s", rec)
		err = identity_provider_error.WrapString("Unexpected error type")
	}

	identityProviderError, isIdentityProviderError := err.(identity_provider_error.IdentityProviderError)
	if !isIdentityProviderError {
		identityProviderError = identity_provider_error.WrapError(err)
	}

	log.Printf(
		"%s (%s) %s",
		identityProviderError.Code(),
		identityProviderError.Title(),
		identityProviderError.Error(),
	)
	log.Print(err)

	body, _ := json.Marshal(dto.ErrorEntryDTO{
		Code: identityProviderError.Code(),
		Title: identityProviderError.Title(),
		Message: identityProviderError.Error(),
	})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(identityProviderError.Status())
	w.Write(body)
}