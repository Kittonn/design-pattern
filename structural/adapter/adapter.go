package adapter

type Storage interface {
	Download(filename string) ([]byte, error)
}

func ProcessStorage(s Storage, filename string) ([]byte, error) {
	return s.Download(filename)
}
