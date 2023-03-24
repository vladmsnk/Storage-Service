package repository

import (
	"git.miem.hse.ru/1206/material-library/internal/client"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

// StorageRepository repository layer with S3 connection
type StorageRepository struct {
	storage *client.StorageClient
}

func NewStorageRepository(storage *client.StorageClient) *StorageRepository {
	return &StorageRepository{storage: storage}
}

func (s *StorageRepository) UploadMaterial(material *model.UploadMaterial) (int64, error) {

	return 0, nil
}
