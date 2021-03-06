package scanlog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})

	With(args ...interface{}) Logger
}

type log struct {
	logger *zap.SugaredLogger
}

func simpleTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func NewLogger(debug bool) (Logger, error) {
	dyn := zap.NewAtomicLevel()

	if debug {
		dyn.SetLevel(zap.DebugLevel)
	} else {
		dyn.SetLevel(zap.InfoLevel)
	}

	config := zap.Config{
		Level:             dyn,
		Encoding:          "console",
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			MessageKey:     "M",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     simpleTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &log{
		logger: logger.Sugar(),
	}, nil
}

func (l *log) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

func (l *log) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *log) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

func (l *log) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

func (l *log) With(args ...interface{}) Logger {
	return &log{
		logger: l.logger.With(args...),
	}
}

func NewNopLogger() Logger {
	return &nop{}
}

type nop struct{}

func (n *nop) Debugf(msg string, args ...interface{}) {
}

func (n *nop) Infof(msg string, args ...interface{}) {
}

func (n *nop) Warnf(msg string, args ...interface{}) {
}

func (n *nop) Errorf(msg string, args ...interface{}) {
}

func (n *nop) With(args ...interface{}) Logger {
	return n
}
