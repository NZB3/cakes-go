package models

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Photo    string `json:"photo,omitempty"`
}
