package logger

import (
	"log"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap() *otelzap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	devLogger, err := config.Build()
	if err != nil {
		log.Fatal("failed to init zap logger")
	}

	logger := otelzap.New(devLogger)

	otelzap.ReplaceGlobals(logger)

	sugar := logger.Sugar()

	return sugar
}
