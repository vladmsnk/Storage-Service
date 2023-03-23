package repository

import (
	"git.miem.hse.ru/1206/app/storage/s3"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

// StorageRepository repository layer with S3 connection
type StorageRepository struct {
	storage s3.Provider
}

func NewStorageRepository() *StorageRepository {
	return &StorageRepository{storage: s3.GetInstance()}
}

func (s *StorageRepository) UploadMaterial(material *model.UploadMaterial) (int64, error) {
	return 0, nil
}
