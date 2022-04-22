package models

import "time"

type ParamSignUp struct {
	Username   string `json:"user_name" binding:"required"`
	Phone      string `json:"phone" binding:"required,len=11"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type User struct {
	UserID     int64
	PhoneValid int
	EmailValid int
	Gender     int
	Status     int
	Username   string
	Password   string
	Phone      string
	Email      string
	SignUpTime time.Time
	LastActive time.Time
}
