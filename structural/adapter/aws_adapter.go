package adapter

type AWSAdapter struct {
	client *AWSS3Client
}

func NewAWSAdapter(client *AWSS3Client) *AWSAdapter {
	return &AWSAdapter{client: client}
}

func (a *AWSAdapter) Download(filename string) ([]byte, error) {
	return a.client.GetObject(filename)
}
