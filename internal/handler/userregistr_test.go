package handler

import (
	"DictionaryComparator/internal/model"
	"DictionaryComparator/internal/repository"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mockDb := new(MockDB)
	user := model.User{Email: "example@email.com", Password: "hashedpassword"}
	mockDb.On("Create", &user).Return(&gorm.DB{})
	err := repository.CreateUser(mockDb, &user)
	if err != nil {
		t.Errorf("CreateUser returned error: %v", err)
	}
	mockDb.AssertExpectations(t)
}
