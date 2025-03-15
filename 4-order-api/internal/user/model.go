package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string
	SessionId   string
	Code        string
}
