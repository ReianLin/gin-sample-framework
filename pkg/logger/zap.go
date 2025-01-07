package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
	logger *zap.SugaredLogger
	writer []io.Writer
}

func NewZapLogger(level zapcore.Level) *Zap {
	var (
		writers = make([]io.Writer, 0, 0)
	)

	writers = append(writers, os.Stderr)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	w := zapcore.AddSync(io.MultiWriter(writers...))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level)
	z := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Zap{
		logger: z.Sugar(),
		writer: writers,
	}
}

func (z *Zap) WithContext(ctx context.Context) Logger {
	newZap := &Zap{
		logger: z.logger,
		writer: z.writer,
	}
	return newZap
}

func (z *Zap) WithField(key string, value interface{}) Logger {
	field := zap.Any(key, value)
	newLogger := z.logger.With(field)
	newZop := &Zap{
		logger: newLogger,
		writer: z.writer,
	}
	return newZop
}

func (z *Zap) WithFields(fields Fields) Logger {
	zapFields := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	newLogger := z.logger.With(zapFields...)

	newZop := &Zap{
		logger: newLogger,
		writer: z.writer,
	}
	return newZop
}

func (z *Zap) Debug(args ...interface{}) {
	z.logger.Debug(args...)
}

func (z *Zap) Debugf(format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

func (z *Zap) Info(args ...interface{}) {
	z.logger.Info(args...)
}

func (z *Zap) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z *Zap) Warn(args ...interface{}) {
	z.logger.Warn(args...)
}

func (z *Zap) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z *Zap) Error(args ...interface{}) {
	z.logger.Error(args...)
}

func (z *Zap) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}

func (z *Zap) Panic(args ...interface{}) {
	z.logger.Panic(args)
}

func (z *Zap) Panicf(format string, args ...interface{}) {
	z.logger.Panicf(format, args...)
}

func (z *Zap) Fatal(args ...interface{}) {
	z.logger.Fatal(args...)
}

func (z *Zap) Fatalf(format string, args ...interface{}) {
	z.logger.Fatalf(format, args...)
}
