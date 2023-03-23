package cmd

import (
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/s3"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"git.miem.hse.ru/1206/material-library/internal/domain"
	"git.miem.hse.ru/1206/material-library/internal/domain/repository"
	"git.miem.hse.ru/1206/material-library/internal/service"
)

func NewDefault(cfg *config.Config) *Cmd {
	c := &Cmd{}
	logger.Init(&cfg.Logger)

	err := s3.InitConnect(&cfg.Storage)
	if err != nil {
		logger.Get().Fatal(err)
	}
	useCase := domain.NewUseCase(repository.NewStorageRepository(), repository.NewMaterialInfoRepo())

	c.grpcServer, err = service.NewLibraryServer(&cfg.GRPC, *useCase)
	if err != nil {
		logger.Get().Fatal(err)
	}
	logger.Get().Info("start grpc server ", cfg.GRPC.Host, ":", cfg.GRPC.Port)

	return c
}
