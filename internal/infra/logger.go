package infra

import "go.uber.org/zap"

func NewLogger(config *AppConfig) (*zap.Logger, error) {
	if config.Development {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		return logger, nil
	} else {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return logger, nil
	}
}

func NewSugaredLogger(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
