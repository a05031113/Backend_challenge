package models

type Users struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
