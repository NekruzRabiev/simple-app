package domain

type User struct {
	Id       int    `json:"-" db:"id"`
	FullName string `json:"fullName" db:"full_name"`
	Password string `json:"-" db:"password"`
	Email    string `json:"email" db:"email"`
}
