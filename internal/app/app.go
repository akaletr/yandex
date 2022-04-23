package app

import (
	"fmt"
	"github.com/akaletr/yandex/internal/storage"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type database struct {
}

var data map[string]string = map[string]string{}

type App interface {
	Write(url string) (string, error)
	Start() error
}

type server struct {
	storage storage.Storage
	name    string
}

func NewServer() (App, error) {
	return &server{}, nil
}

func (app *server) Start() error {
	http.HandleFunc("/", addURL)
	return http.ListenAndServe(":8080", nil)
}

func addURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		link := fmt.Sprint("/", strconv.Itoa(rand.Int()))
		data[link] = string(body)
		_, _ = w.Write([]byte(link))
	case http.MethodGet:
		link, ok := data[r.URL.String()]
		if ok {
			w.Header().Set("Location", data[link])
			w.WriteHeader(http.StatusTemporaryRedirect)
		}

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (app *server) Write(url string) (string, error) {
	fmt.Println("It works")
	return url, nil
}
