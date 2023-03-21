package logger

import (
	"git.miem.hse.ru/1206/microservice-template/internal/config"
	"go.uber.org/zap"
	"log"
)

// Название этой директории также говорит самой за себя.
// Здесь располагается логгер и метод его инициализации.
// В большинстве случаев для данного пакета используется паттерн Singletone.

var logger *zap.SugaredLogger

func Get() *zap.SugaredLogger {
	return logger
}

func Init(cfg *config.Config) {
	loggerLevel, err := zap.ParseAtomicLevel(cfg.Logger.Level)
	if err != nil {
		log.Fatal("invalid logger level")
	}

	loggerConfig := zap.Config{
		Level:             loggerLevel,
		Development:       cfg.Logger.Development,
		DisableCaller:     !cfg.Logger.Caller,
		DisableStacktrace: false,
		Encoding:          cfg.Logger.Formatter,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{cfg.Logger.Output},
		ErrorOutputPaths:  []string{cfg.Logger.Output},
	}

	loggerRaw, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("error building logger: %s", err.Error())
	}

	logger = loggerRaw.Sugar()
}
