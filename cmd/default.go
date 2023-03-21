package cmd

import (
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/s3"
	"git.miem.hse.ru/1206/microservice-template/internal/config"
)

func NewDefault(cfg *config.Config) *Cmd {
	c := &Cmd{}

	logger.Init(&cfg.Logger)

	//client, err := minio.New(fmt.Sprintf("%s:%s", cfg.Storage.Host, cfg.Storage.Port), &minio.Options{
	//	Creds:  credentials.NewStaticV4(cfg.Storage.AccessKeyId, cfg.Storage.SecretAccessKey, ""),
	//	Secure: cfg.Storage.UseSsl,
	//})
	//fmt.Println(client)
	//if err != nil {
	//	logger.Get().Fatal(err)
	//}

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
