package main

import (
	"DictionaryComparator/internal/handler"
	"DictionaryComparator/internal/local"
	"DictionaryComparator/internal/model"
	"DictionaryComparator/internal/security"
	"log"
	"net/http"
)

func main() {
	//DB starter
	var db local.GormDBInterface
	db, err := local.NewDatabase()
	if err != nil {
		log.Fatal("Could not connect to the database: %v", err)
	}
	//DB migration
	db.AutoMigrate(&model.User{}, &model.Word{})
	//open endpoints
	http.HandleFunc("/create_user", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateUserHandler(db, w, r) // передаем *gorm.DB в обработчик
	})
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		handler.LoginHandler(db, writer, request)
	})
	//secured endpoints
	http.Handle("/get_all_words", security.AuthMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetAllWordsHandler(db, w, r)
		})))
	http.Handle("/check_words", security.AuthMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.CheckWordsHandler(db, w, r)
		})))
	http.Handle("/add_bulk_words", security.AuthMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.AddBulkWordsHandler(db, w, r)
		})))
	//html sheet dependency
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	//server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
