package db

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/joho/godotenv"
)

type BlobStorage struct {
	containerURL azblob.ContainerURL
}

func InitBlobStorage(containerName string) (*BlobStorage, error) {
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

	return &BlobStorage{
		containerURL: containerURL,
	}, nil
}

func (bs *BlobStorage) UploadFile(ctx context.Context, filename string, data io.Reader) (string, error) {
	blobURL := bs.containerURL.NewBlockBlobURL(filename)
	_, err := azblob.UploadStreamToBlockBlob(ctx, data, blobURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		return "", err
	}

	return blobURL.String(), nil
}
