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

func (bs *BlobStorage) UploadFile(ctx context.Context, f models.FileData) (string, error) {
	blobURL := bs.containerURL.NewBlockBlobURL(f.Path)
	_, err := azblob.UploadStreamToBlockBlob(ctx, f.Data, blobURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		return "", err
	}

	return blobURL.String(), nil
}

func (bs *BlobStorage) DownloadFile(ctx context.Context, path string) (*models.FileData, error) {
	blobURL := bs.containerURL.NewBlockBlobURL(path)
	resp, err := blobURL.Download(ctx, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		return nil, err
	}

	f := &models.FileData{
		Path: path,
		Data: resp.Body(azblob.RetryReaderOptions{}),
	}
	return f, nil
}
