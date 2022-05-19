package models

import "time"

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Phone      string `json:"phone" binding:"required,len=11"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserID     int64
	PhoneValid int
	EmailValid int
	Gender     int
	Status     int
	Username   string `db:"username"`
	Password   string `db:"password"`
	Phone      string
	Email      string
	SignUpTime time.Time
	LastActive time.Time
}
