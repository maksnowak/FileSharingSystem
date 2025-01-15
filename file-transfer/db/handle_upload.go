package db

import (
	"context"
	"fmt"
	"os"

	"file-transfer/models"
)

func (bs *LocalBlobStorage) UploadFile(ctx context.Context, f models.FileData) (string, error) {
	path := fmt.Sprintf("%s/%s", bs.rootPath, f.Path)
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write()
	if err != nil {
		return "", err
	}

	return path, nil
}

func (bs *LocalBlobStorage) DownloadFile(ctx context.Context, path string) (*models.FileData, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s", bs.rootPath, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	f := &models.FileData{
		Path: path,
		Data: file,
	}
	return f, nil
}

func (bs *AzureBlobStorage) UploadFile(ctx context.Context, f models.FileData) (string, error) {
	blobURL := bs.containerURL.NewBlockBlobURL(f.Path)
	_, err := azblob.UploadStreamToBlockBlob(ctx, f.Data, blobURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		return "", err
	}

	return blobURL.String(), nil
}

func (bs *AzureBlobStorage) DownloadFile(ctx context.Context, path string) (*models.FileData, error) {
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
