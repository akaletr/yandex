package app

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var data map[string]string = map[string]string{}

type App interface {
	Write(url string) (string, error)
	Start() error
}

type server struct {
}

func NewServer() (App, error) {
	return &server{}, nil
}

func (app *server) Start() error {
	http.HandleFunc("/", addURL)
	return http.ListenAndServe(":8080", nil)
}

func addURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println(data)

	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		hasher := sha1.New()
		hasher.Write(body)
		linkID := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		data[linkID] = string(body)

		link := url.URL{
			Scheme: "http",
			Host:   r.Host,
			Path:   linkID,
		}

		_, _ = w.Write([]byte(link.String()))
	case http.MethodGet:
		link, ok := data[strings.TrimPrefix(r.URL.Path, "/")]

		fmt.Println("linkID", link)
		if ok {
			w.Header().Set("Location", link)
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
