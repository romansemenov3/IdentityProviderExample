package request_handler

import (
	"common"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"model/identity_provider_error"
	"net/http"
	"strings"
)

type config struct {
	Auth authConfig `yaml:"auth"`
}

type authConfig struct {
	Secret string `yaml:"secret"`
}

var secret string

func init() {
	cfg := config{}
	common.ReadConfig(&cfg)

	secret = cfg.Auth.Secret
}

func TokenValidator(inner http.Handler, grants []string) http.Handler {
	if len(grants) == 0 {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			inner.ServeHTTP(w, r)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := verifyToken(r)
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			actualGrants, ok := claims["grants"].([]interface{})
			if !ok {
				panic(identity_provider_error.ForbiddenError{})
			}
			grantsMap := map[interface{}]bool{}
			for _, grant := range actualGrants {
				grantsMap[grant] = true
			}
			for _, grant := range grants {
				if _, hasKey := grantsMap[grant]; !hasKey {
					panic(identity_provider_error.ForbiddenError{})
				}
			}
		}

		inner.ServeHTTP(w, r)
	})
}

func verifyToken(r *http.Request) *jwt.Token {
	tokenString := extractToken(r)
	if tokenString == "" {
		panic(identity_provider_error.UnauthorizedError{})
	}

	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		log.Print(err)
		panic(identity_provider_error.UnauthorizedError{})
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		panic(identity_provider_error.UnauthorizedError{})
	}
	return token
}

func extractToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authorization, "Bearer ")
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(secret), nil
}
