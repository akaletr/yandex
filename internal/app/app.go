package app

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/akaletr/yandex/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App interface {
	Start() error
}

type server struct {
	storage storage.Storage
}

func NewServer() (*server, error) {
	return &server{
		storage: storage.NewStorage(),
	}, nil
}

func (app *server) Start() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/", app.addLink)
	router.Get("/{id}", app.getLink)

	return http.ListenAndServe(":8080", router)
}

func (app *server) addLink(w http.ResponseWriter, r *http.Request) {
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
}

func (app *server) getLink(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	link, err := app.storage.Read(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(""))
	}

	w.Header().Set("Location", link)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
