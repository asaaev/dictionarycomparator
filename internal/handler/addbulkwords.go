package handler

import (
	"DictionaryComparator/internal/local"
	"DictionaryComparator/internal/model"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func AddBulkWordsHandler(db local.GormDBInterface, w http.ResponseWriter, r *http.Request) {
	var request struct {
		Words string `json:"words"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wordsToAdd := strings.Split(request.Words, ",")

	//status tracing map
	addStatus := make(map[string]string)
	for _, word := range wordsToAdd {
		normalizedWord := strings.ToLower(strings.TrimSpace(word))
		//empty word check
		if normalizedWord == "" {
			addStatus[word] = "ignored - empty word"
		}
		var wordModel model.Word
		if err := db.Where("word_body = ?", normalizedWord).First(&wordModel).Error; gorm.ErrRecordNotFound == nil {
			//if word already exists
			addStatus[normalizedWord] = "ignored - already exists"
			continue
		} else if err != nil && err != gorm.ErrRecordNotFound {
			addStatus[normalizedWord] = "error - during lookup"
			continue
		}
		//if word not present
		newWord := model.Word{WordBody: normalizedWord}
		if err := db.Create(&newWord).Error; err != nil {
			//error during injection
			addStatus[normalizedWord] = "error - during add"
			continue
		}
		addStatus[normalizedWord] = "added"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addStatus)
}
