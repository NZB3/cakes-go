package models

type User struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role,omitempty" db:"role"`
	FullName string `json:"full_name,omitempty" db:"full_name"`
	Photo    string `json:"photo,omitempty" db:"photo"`
}
