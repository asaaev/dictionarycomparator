package handler

import (
	"DictionaryComparator/internal/model"
	"DictionaryComparator/internal/security"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

// LoginRequest receive data from user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//initialize structure
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//find user by email & check pass
	var user model.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "invalid email login credentials", http.StatusUnauthorized)
		return
	}
	if err := user.CheckPassword(req.Password); err != nil {
		http.Error(w, "invalid password login credentials", http.StatusUnauthorized)
		return
	}

	token, err := security.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "error while generate token", http.StatusInternalServerError)
		return
	}
	//send jwt to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
