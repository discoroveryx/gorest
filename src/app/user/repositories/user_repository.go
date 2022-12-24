package repositories

import (
	"app/user/models"
	"dbstorage"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db := new(dbstorage.DB)
	cursor := db.Connect()
	return UserRepository{db: cursor}
}

func (u *UserRepository) UserExistsByName(name string) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.db.Where("name = ?", name).Take(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) GetUserById(id uint) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.db.Where("id = ?", id).First(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) UserExistsByEmail(email string) bool {
	result := u.db.Where("email = ?", email).First(&models.UserModel{})

	// fmt.Println(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (u *UserRepository) UpdateVerificationCode(user models.UserModel, code string) (models.UserModel, error) {
	user.VerificationCode = code
	result := u.db.Save(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) IsUserVerifiedById(id uint) (bool, error) {
	user := models.UserModel{}

	result := u.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return false, result.Error
	}
	// fmt.Println(result.Error)

	return user.Verified, result.Error
}

func (u *UserRepository) UserVerify(userId uint, verificationCode string) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.db.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return user, result.Error
	}
	// fmt.Println(result.Error)

	user.Verified = true
	result = u.db.Save(&user)

	return user, result.Error
}
