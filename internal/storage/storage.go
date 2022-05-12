package storage

import (
	"errors"
)

type storage struct {
	data map[string]string
}

func NewStorage() Storage {
	return &storage{
		data: map[string]string{},
	}
}

func (storage *storage) Write(key, value string) error {
	storage.data[key] = value
	return nil
}

func (storage *storage) Read(id string) (string, error) {
	link, ok := storage.data[id]
	if ok {
		return link, nil
	}

	return "", errors.New("not found")
}
