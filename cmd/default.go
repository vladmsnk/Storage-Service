package cmd

import (
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/s3"
	"git.miem.hse.ru/1206/material-library/internal/config"
)

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
