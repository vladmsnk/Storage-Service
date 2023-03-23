package repository

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"github.com/minio/minio-go/v7"
)

// StorageRepository repository layer with S3 connection
type StorageRepository struct {
	client *minio.Client //s3 instance
}

func NewStorageRepository(client *minio.Client) *StorageRepository {
	return &StorageRepository{client: client}
}

func (s *StorageRepository) UploadMaterial(material *model.UploadMaterial) (int64, error) {
	return 0, nil
}
