package logger

import "go.uber.org/zap"

type Level string

var (
	Dev  Level = "dev"
	Prod Level = "prod"
)

func New(level Level) *zap.SugaredLogger {
	var (
		logger = new(zap.Logger)
		err    error
	)
	switch level {
	case "prod":
		logger, err = zap.NewProduction()
		if err != nil {
			return nil
		}
	case "dev":
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil
		}
	}
	return logger.Sugar()
}
