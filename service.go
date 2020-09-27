package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Service interface {
	GetShortUrl(w http.ResponseWriter, r *http.Request)
	RedirectByUrl(w http.ResponseWriter, r *http.Request)
}

func NewService(store Store) Service {
	return &service{store: store}
}

type service struct {
	store Store
}

func (s service) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params)
}

func (s service) RedirectByUrl(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}