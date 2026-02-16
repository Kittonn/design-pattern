package adapter

type AWSS3Client struct{}

func (a *AWSS3Client) GetObject(key string) ([]byte, error) {
	return []byte("data from AWS S3"), nil
}
