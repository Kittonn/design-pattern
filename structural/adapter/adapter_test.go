package adapter

import "testing"

func TestGCPStorage(t *testing.T) {
	gcpClient := &GCPStorageClient{}
	gcpAdapter := NewGCPAdapter(gcpClient)

	data, err := gcpAdapter.Download("file.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "data from GCP Storage"
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestAWSS3Storage(t *testing.T) {
	awsClient := &AWSS3Client{}
	awsAdapter := NewAWSAdapter(awsClient)

	data, err := awsAdapter.Download("file.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "data from AWS S3"
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}
