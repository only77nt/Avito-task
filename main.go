package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/only77nt/avito-task/adapter"
	service2 "github.com/only77nt/avito-task/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	//dsn := "root:QWEasd777+@tcp(127.0.0.1:3306)/links?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(172.18.0.2:3306)/links?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	store := adapter.NewGormStore(db)

	service := service2.NewService(store)

	r := mux.NewRouter()

	r.HandleFunc("/short-url", service.GetShortUrl)
	r.HandleFunc("/redirect", service.RedirectByUrl)

	fmt.Println("Start listening in http://localhost:9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}