package db

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"file-transfer/models"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/joho/godotenv"
)

type BlobStorage interface {
	UploadFile(ctx context.Context, f models.FileData) (string, error)
	DownloadFile(ctx context.Context, path string) (*models.FileData, error)
}

type LocalBlobStorage struct {
	rootPath string
}

func InitLocalBlobStorage(rootPath string) (*LocalBlobStorage, error) {
	return &LocalBlobStorage{
		rootPath: rootPath,
	}, nil
}

type AzureBlobStorage struct {
	containerURL azblob.ContainerURL
}

func InitAzureBlobStorage(containerName string) (*AzureBlobStorage, error) {
	_ = godotenv.Load("./.env")

	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %v", err)
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	containerURL := azblob.NewContainerURL(*URL, pipeline)

	return &AzureBlobStorage{
		containerURL: containerURL,
	}, nil
}
