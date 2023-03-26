package repository

import (
	"bytes"
	"context"
	"git.miem.hse.ru/1206/material-library/internal/client"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"github.com/minio/minio-go/v7"
	"io"
)

// StorageRepository repository layer with S3 connection
type StorageRepository struct {
	storage *client.StorageClient
	config  *config.Storage
}

func NewStorageRepository(storage *client.StorageClient, config *config.Storage) *StorageRepository {
	return &StorageRepository{storage: storage, config: config}
}

func (s *StorageRepository) UploadMaterial(m *model.Material) (string, error) {

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

func (s *StorageRepository) GetMaterialByObjectName(objectName string) (*model.Material, error) {
	object, err := s.storage.Client.GetObject(context.Background(), s.config.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return &model.Material{Reader: bytes.NewReader(data)}, nil
}

func (s *StorageRepository) CheckMaterialExists(objectName string) (bool, error) {
	_, err := s.storage.Client.StatObject(context.Background(), s.config.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

//func (s *StorageRepository) CheckMaterialExistsByObjectName(objectName string) (bool, error) {
//	return s.storage.Client.
//
//}
