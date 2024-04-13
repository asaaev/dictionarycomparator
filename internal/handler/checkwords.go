package handler

import (
	"DictionaryComparator/internal/model"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func CheckWordsHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var request struct {
		Words string `json:"words"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wordsToCheck := strings.Split(request.Words, ",")
	existingWords := make(map[string]bool)
	for _, word := range wordsToCheck {
		var wordModel model.Word
		if err := db.Where("word_body = ?", strings.TrimSpace(word)).First(&wordModel).Error; err == nil {
			existingWords[word] = true
		} else {
			existingWords[word] = false
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingWords)
}
