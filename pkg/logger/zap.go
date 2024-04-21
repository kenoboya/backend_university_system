package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(lvl zapcore.Level, args ...interface{})
	Logf(lvl zapcore.Level, format string, args ...interface{})
}

func CreateLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
	return zap.Must(config.Build())
}

func Debug(args ...interface{}) {
	zap.S().Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}
func Info(args ...interface{}) {
	zap.S().Info(args...)
}
func Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}
func Warn(args ...interface{}) {
	zap.S().Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}
func Error(args ...interface{}) {
	zap.S().Error(args...)
}
func Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}
func Panic(args ...interface{}) {
	zap.S().Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	zap.S().Panicf(format, args...)
}
func Fatal(args ...interface{}) {
	zap.S().Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(format, args...)
}
func Log(lvl zapcore.Level, args ...interface{}) {
	zap.S().Log(lvl, args...)
}
func Logf(lvl zapcore.Level, format string, args ...interface{}) {
	zap.S().Logf(lvl, format, args...)
}
