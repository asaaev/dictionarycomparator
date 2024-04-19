package handler

import (
	"DictionaryComparator/internal/local"
	"DictionaryComparator/internal/model"
	"encoding/json"
	"net/http"
)

func GetAllWordsHandler(db local.GormDBInterface, w http.ResponseWriter, r *http.Request) {
	var words []model.Word
	if err := db.Order("word_body asc").Find(&words).Error; err != nil {
		http.Error(w, "Error retrieving words", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}
