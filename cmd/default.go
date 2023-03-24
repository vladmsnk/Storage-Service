package cmd

import (
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/material-library/internal/client"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"git.miem.hse.ru/1206/material-library/internal/domain"
	"git.miem.hse.ru/1206/material-library/internal/domain/repository"
	"git.miem.hse.ru/1206/material-library/internal/service"
)

func NewDefault(cfg *config.Config) *Cmd {
	c := &Cmd{}
	logger.Init(&cfg.Logger)

	cl, err := client.NewStorageClient(&cfg.Storage)
	if err != nil {
		logger.Get().Fatal(err)
	}
	useCase := domain.NewUseCase(repository.NewStorageRepository(cl), repository.NewMaterialInfoRepo())

	c.grpcServer, err = service.NewLibraryServer(&cfg.GRPC, *useCase)
	if err != nil {
		logger.Get().Fatal(err)
	}
	logger.Get().Info("start grpc server ", cfg.GRPC.Host, ":", cfg.GRPC.Port)

	return c
}
