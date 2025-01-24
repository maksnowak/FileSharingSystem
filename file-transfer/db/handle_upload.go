package db

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"file-transfer/models"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func (bs *LocalBlobStorage) UploadFile(ctx context.Context, f models.FileData) (string, error) {
  path := fmt.Sprintf("%s/%s/%s/%s", "http://localhost:8080", bs.rootPath, f.UserID, f.Path)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(f.Data)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (bs *LocalBlobStorage) DownloadFile(ctx context.Context, userID string, path string) (*models.FileData, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s", bs.rootPath, userID, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := io.Reader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	f := &models.FileData{
		Path: path,
		Data: data,
	}
	return f, nil
}

func (bs *AzureBlobStorage) UploadFile(ctx context.Context, f models.FileData) (string, error) {
	reader := bytes.NewReader(f.Data)

	blobURL := bs.containerURL.NewBlockBlobURL(fmt.Sprintf("%s/%s", f.UserID, f.Path))
	_, err := azblob.UploadStreamToBlockBlob(ctx, reader, blobURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		return "", err
	}

	return blobURL.String(), nil
}

func (bs *AzureBlobStorage) DownloadFile(ctx context.Context, userID string, path string) (*models.FileData, error) {
	blobURL := bs.containerURL.NewBlockBlobURL(fmt.Sprintf("%s/%s", userID, path))
	resp, err := blobURL.Download(ctx, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		return nil, err
	}

	reader := resp.Body(azblob.RetryReaderOptions{})
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	f := &models.FileData{
		Path: path,
		Data: data,
	}
	return f, nil
}
