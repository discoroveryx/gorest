package repositories

import (
	"app/user/models"
	"dbstorage"

	"gorm.io/gorm"
)

type UserCreateRepository struct {
	db *gorm.DB
}

func NewUserCreateRepository() UserCreateRepository {
	db := new(dbstorage.DB)
	cursor := db.Connect()
	return UserCreateRepository{db: cursor}
}

func (u *UserCreateRepository) UserCreate(name string, email string, password string) models.UserModel {
	user := models.UserModel{
		Name:     name,
		Email:    email,
		Password: password,
		Verified: false,
		// Ctime:    time.Now(),
	}

	u.db.Create(&user)
	return user
}
