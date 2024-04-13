package repository

import (
	"DictionaryComparator/internal/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}
func GetUserByEmail(db *gorm.DB, email string) (model.User, error) {
	var user model.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}
