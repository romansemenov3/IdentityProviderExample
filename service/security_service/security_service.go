package security_service

import (
	"common"
	"github.com/dgrijalva/jwt-go"
	"model/dto"
	"model/entity"
	"model/identity_provider_error"
	"service/client_service"
	"service/grant_service"
	"time"
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

func GenerateToken(user entity.UserEntity, clientId string) dto.TokenDTO {
	var key []byte
	if clientId == "" {
		key = []byte(secret)
	} else {
		client := client_service.GetByIdOrThrow(clientId)
		if client == nil {
			panic(identity_provider_error.ClientNotFoundError{Id: clientId})
		}
		key = []byte(client.Secret)
	}

	grants := grant_service.GetGrantsByUserId(user.Id)
	expiresAt := time.Now().Add(time.Millisecond * 7200000).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["user"] = user.Id.String()
	atClaims["grants"] = grants
	atClaims["exp"] = expiresAt
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString(key)
	if err != nil {
		panic(err)
	}

	return dto.TokenDTO{
		AccessToken: accessToken,
		TokenType:   "bearer",
		ExpiresIn:   7200000,
	}
}
