package logger

import (
	"go.uber.org/zap"
)

type Stage string

var (
	Dev  Stage = "dev"
	Prod Stage = "prod"
)

func New(stage Stage, logLevel string) (*zap.SugaredLogger, error) {
	var (
		logger = new(zap.Logger)
		err    error
	)
	switch stage {
	case Prod:
		if logger, err = zap.NewProduction(); err != nil {
			return nil, err
		}
	case Dev:
		if logger, err = zap.NewDevelopment(); err != nil {
			return nil, err
		}
	}
	level := logger.Level()
	if err = level.Set(logLevel); err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
