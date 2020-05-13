package grant_service

import (
	"github.com/gofrs/uuid"
	"service/grant_repository"
)

func GetGrantsByUserId(userId uuid.UUID) []string {
	grants := grant_repository.FindByUserId(userId)
	var grantsCodes = make([]string, len(grants))
	for i, grant := range grants {
		grantsCodes[i] = grant.Code
	}
	return grantsCodes
}
