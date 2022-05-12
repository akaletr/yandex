package storage

type Storage interface {
	Write(key, value string) error
	Read(key string) (string, error)
}
