package handler

import (
	"DictionaryComparator/internal/model"
	"DictionaryComparator/internal/repository"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func CreateUserHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if user already exist with that email
	var existingUser model.User
	if err := db.Where("email = ?", newUser.Email).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	// If user doesn't exist, hashing and adding
	if err := newUser.HashPassword(newUser.Password); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	if err := repository.CreateUser(db, &newUser); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
