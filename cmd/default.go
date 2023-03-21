package cmd

import (
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/s3"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	client *minio.Client
}

var instance *Storage

func NewDefault(cfg *config.Config) *Cmd {
	c := &Cmd{}
	logger.Init(&cfg.Logger)

	err := s3.InitConnect(&cfg.Storage)
	if err != nil {
		logger.Get().Fatal(err)
	}

	// добавление всех зависимостей для запуска приложения в обычном режиме

	return c
}

//func InitConnect(cfg *config.) error {
//	var err error
//	instance.client, err = minio.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), &minio.Options{
//		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKey, ""),
//		Secure: cfg.UseSsl,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
