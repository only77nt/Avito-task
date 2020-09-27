package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/only77nt/avito-task/adapter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "root:QWEasd777+@tcp(127.0.0.1:3306)/links?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	store := adapter.NewGormStore(db)

	service := NewService(store)

	r := mux.NewRouter()

	r.HandleFunc("/{url}", service.GetShortUrl).Methods("GET")
	r.HandleFunc("/", service.RedirectByUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}