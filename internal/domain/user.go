package domain

type User struct {
	Id       int    `json:"id,omitempty" db:"id"`
	FullName string `json:"fullName,omitempty" db:"full_name"`
	Password string `json:"-" db:"password"`
	Email    string `json:"email,omitempty" db:"email"`
}
