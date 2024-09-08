package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		var err error
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout", "./app.log"}
		logger, err = config.Build()
		if err != nil {
			panic("Error init zap logger" + err.Error())
		}
	})
}

func Global() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

func CheckError(data string, err error) error {
	if err != nil {
		logger.Error(data, zap.Error(err))
		return err
	}

	return nil
}

func CheckInfo(data string) {
	if data != "" {
		logger.Info(data)
	}
}
