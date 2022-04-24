package storage

import (
	"crypto/sha1"
	"encoding/base64"
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

func (storage *storage) Write(link string) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(link))
	if err != nil {
		return "", err
	}

	linkID := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	storage.data[linkID] = link

	return linkID, nil
}

func (storage *storage) Read(id string) (string, error) {
	link, ok := storage.data[id]
	if ok {
		return link, nil
	}

	return "", errors.New("not found")
}
