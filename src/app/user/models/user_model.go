package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name             string
	Email            string
	Password         string
	VerificationCode string
	Verified	bool
	// Ctime    time.Time
}
