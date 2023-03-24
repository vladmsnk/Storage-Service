package client

import (
	"fmt"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageClient struct {
	Client *minio.Client
}

func NewStorageClient(cfg *config.Storage) (*StorageClient, error) {
	var err error
	options := &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSsl,
	}
	s3Client, err := minio.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), options)
	if err != nil {
		return nil, err
	}
	return &StorageClient{Client: s3Client}, err
}
