package dto

type UserDTO struct {
	Id       string   `json:"id,omitempty"`
	Grants   []string `json:"grants,omitempty"`
	Login    string   `json:"login,omitempty"`
	Password string   `json:"password,omitempty"`
}
