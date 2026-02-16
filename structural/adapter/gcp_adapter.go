package adapter

type GCPAdapter struct {
	client *GCPStorageClient
}

func NewGCPAdapter(client *GCPStorageClient) *GCPAdapter {
	return &GCPAdapter{client: client}
}

func (a *GCPAdapter) Download(filename string) ([]byte, error) {
	return a.client.ReadObject(filename)
}
