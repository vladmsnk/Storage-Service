package repository

import (
	"context"
	"git.miem.hse.ru/1206/material-library/internal/client"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"github.com/minio/minio-go/v7"
)

// StorageRepository repository layer with S3 connection
type StorageRepository struct {
	storage *client.StorageClient
	config  *config.Storage
}

func NewStorageRepository(storage *client.StorageClient, config *config.Storage) *StorageRepository {
	return &StorageRepository{storage: storage, config: config}
}

func (s *StorageRepository) UploadMaterial(m *model.UploadMaterial) (string, error) {

	uploadInfo, err := s.storage.Client.PutObject(context.Background(), s.config.BucketName, m.ObjectName, m.Reader,
		m.ObjectSize,
		minio.PutObjectOptions{ContentType: m.ContentType})
	if err != nil {
		return "", err
	}
	return uploadInfo.Key, nil
}

func (s *StorageRepository) DeleteMaterialByObjectName(objectName string) error {
	err := s.storage.Client.RemoveObject(context.Background(), s.config.BucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
