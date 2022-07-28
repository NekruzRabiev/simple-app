package domain

type User struct {
	Id       int    `json:"-" db:"id"`
	FullName string `json:"fullName" db:"full_name"`
	Password string `json:"-" db:"password"`
	Phone    string `json:"phone" db:"phone"`
}
