package adapter

type GCPStorageClient struct{}

func (g *GCPStorageClient) ReadObject(objectName string) ([]byte, error) {
	return []byte("data from GCP Storage"), nil
}
