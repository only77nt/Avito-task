package service

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
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
	query := r.URL.Query()
	url := query["url"][0]

	if url == "" {
		fmt.Println("Empty url")
	}

	var shortUrl string
	if query["need"] != nil {
		shortUrl = query["need"][0]
	} else {
		shortUrl = guuid.New().String()
	}

	if err := s.store.SaveUrlInDB(url, shortUrl); err != nil {
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode(&shortUrl); err != nil {
		fmt.Println(err)
	}
}

func (s service) RedirectByUrl(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	shortUrl := query["url"][0]

	if shortUrl == "" {
		fmt.Println("Empty url")
	}

	url, err := s.store.FindUrlInDB(shortUrl)
	if err != nil {
		fmt.Println(err)
	}

	*url = "https://" + *url

	http.Redirect(w, r, *url, http.StatusTemporaryRedirect)
}