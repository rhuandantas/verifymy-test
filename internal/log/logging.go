package log

import (
	"fmt"
	"github.com/rhuandantas/verifymy-test/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//go:generate mockgen -source=$GOFILE -package=mock_log -destination=../../test/mock/log/$GOFILE

type SimpleLogger interface {
	Debugf(template string, args ...interface{})
	Debug(args ...interface{})
	Infof(template string, args ...interface{})
	Info(args ...interface{})
	Errorf(template string, args ...interface{})
	Error(args ...interface{})
	Warnf(template string, args ...interface{})
	Warn(args ...interface{})
}

type SimpleLoggerImpl struct {
	base *zap.SugaredLogger
}

func (il *SimpleLoggerImpl) Debugf(template string, args ...interface{}) {
	il.base.Debugf(template, args...)
}

func (il *SimpleLoggerImpl) Debug(args ...interface{}) {
	il.base.Debug(args...)
}

func (il *SimpleLoggerImpl) Infof(template string, args ...interface{}) {
	il.base.Infof(template, args...)
}

func (il *SimpleLoggerImpl) Info(args ...interface{}) {
	il.base.Info(args...)
}

func (il *SimpleLoggerImpl) Errorf(template string, args ...interface{}) {
	il.base.Errorf(template, args...)
}

func (il *SimpleLoggerImpl) Error(args ...interface{}) {
	il.base.Error(args...)
}

func (il *SimpleLoggerImpl) Warnf(template string, args ...interface{}) {
	il.base.Warnf(template, args...)
}

func (il *SimpleLoggerImpl) Warn(args ...interface{}) {
	il.base.Warn(args...)
}

func NewLogger(configStore config.ConfigProvider) SimpleLogger {
	zapLevel := zap.NewAtomicLevel()
	logLevel := configStore.GetString("log.level")
	err := zapLevel.UnmarshalText([]byte(logLevel))
	if err != nil {
		fmt.Printf("can't configure logger - %v\n", err)
	}

	logger, _ := zap.Config{
		Encoding:         "json",
		Level:            zapLevel,
		OutputPaths:      []string{os.Stdout.Name()},
		ErrorOutputPaths: []string{os.Stderr.Name()},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()

	defer func() { // flushes buffer, if any
		err := logger.Sync()
		if err != nil {
			fmt.Printf("can't flush logger - %v\n", err)
		}
	}()
	sugar := logger.WithOptions(
		zap.AddCallerSkip(1),
	).Sugar().With(
		"version", configStore.GetString("app.version"),
		"app", configStore.GetString("app.name"),
	)

	return &SimpleLoggerImpl{
		base: sugar,
	}
}
