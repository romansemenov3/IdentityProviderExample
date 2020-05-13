package user_service

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"math/rand"
	"model/dto"
	"model/entity"
	"model/identity_provider_error"
	"service/grant_service"
	"service/security_service"
	"service/user_repository"
	"time"
)

func randomHex() string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func GetUserByIdOrThrow(id string) dto.UserDTO {
	uuid, err := uuid.FromString(id)
	if err != nil {
		panic(err)
	}
	user := user_repository.FindById(uuid)
	if user == nil {
		panic(identity_provider_error.UserNotFoundError{Id: id})
	}

	grants := grant_service.GetGrantsByUserId(user.Id)

	return dto.UserDTO{
		Id:     user.Id.String(),
		Grants: grants,
		Login:  user.Login,
	}
}

func DeleteUserById(id string) {
	uuid, err := uuid.FromString(id)
	if err != nil {
		panic(err)
	}
	user_repository.DeleteById(uuid)
}

func CreateUser(userDTO dto.UserDTO) dto.UserDTO {
	salt := randomHex()
	hash := sha256.New()
	hash.Write([]byte(userDTO.Password + salt))
	digest := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	user := entity.UserEntity{
		Login:    userDTO.Login,
		Salt:     salt,
		Password: digest,
	}
	user = user_repository.Create(user)

	grants := grant_service.GetGrantsByUserId(user.Id)

	return dto.UserDTO{
		Id:     user.Id.String(),
		Grants: grants,
		Login:  user.Login,
	}
}

func Login(loginDTO dto.FormLoginDTO) dto.TokenDTO {
	user := user_repository.FindByLogin(loginDTO.Login)
	if user == nil {
		panic(identity_provider_error.UnauthorizedError{})
	}

	hash := sha256.New()
	hash.Write([]byte(loginDTO.Password + user.Salt))
	digest := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	if digest != user.Password {
		panic(identity_provider_error.UnauthorizedError{})
	}

	return security_service.GenerateToken(*user, loginDTO.ClientId)
}
