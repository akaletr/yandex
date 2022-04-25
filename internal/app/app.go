package app

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/akaletr/yandex/internal/storage"
)

type App interface {
	Start() error
}

type server struct {
	storage storage.Storage
}

func NewServer() (App, error) {
	return &server{
		storage: storage.NewStorage(),
	}, nil
}

func (app *server) Start() error {
	http.HandleFunc("/", app.handler)
	return http.ListenAndServe(":8080", nil)
}

func (app *server) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		hasher := sha1.New()
		_, err = hasher.Write(body)

		linkID := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		err = app.storage.Write(linkID, string(body))
		if err != nil {
			fmt.Println(err)
		}

		link := url.URL{
			Scheme: "http",
			Host:   r.Host,
			Path:   linkID,
		}

		_, _ = w.Write([]byte(link.String()))
	case http.MethodGet:
		link, err := app.storage.Read(strings.TrimPrefix(r.URL.Path, "/"))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(""))
		}
		w.Header().Set("Location", link)
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}
