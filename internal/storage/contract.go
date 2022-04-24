package storage

type Storage interface {
	Write(url string) (string, error)
	Read(id string) (string, error)
}
