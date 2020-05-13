package dto

type FormLoginDTO struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	ClientId string `json:"client_id,omitempty"`
}
