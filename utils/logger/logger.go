package logger

import (
	"os"
	"regexp"
	"strings"

	"github.com/ride-app/marketplace-service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	// Trace(args ...interface{})
	// Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]string) Logger
	WithError(err error) Logger
}

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func New(config *config.Config) *ZapLogger {
	encoderConfig := zap.NewProductionEncoderConfig()

	if !config.Production {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "severity"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	if !config.Production {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       !config.Production,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
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

	if !config.Production {
		zapConfig.Encoding = "console"
	}

	if config.LogDebug {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger := zap.Must(zapConfig.Build(zap.AddCallerSkip(1))).Sugar()

	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.logger.Debug(formatString(args...)...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, formatString(args...)...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.logger.Info(formatString(args...)...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, formatString(args...)...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(formatString(args...)...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, formatString(args...)...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.logger.Error(formatString(args...)...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, formatString(args...)...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(formatString(args...)...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, formatString(args...)...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(formatString(args...)...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, formatString(args...)...)
}

func (l *ZapLogger) WithField(key string, value interface{}) Logger {
	return &ZapLogger{
		logger: l.logger.With(key, value),
	}
}

func (l *ZapLogger) WithFields(fields map[string]string) Logger {
	logger := l.logger
	for key, value := range fields {
		logger = logger.With(key, value)
	}
	return &ZapLogger{
		logger: logger,
	}
}

func formatString(args ...interface{}) []interface{} {
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	formattedArgs := make([]interface{}, len(args))
	for i, arg := range args {
		str, ok := arg.(string)
		if ok {
			str = re.ReplaceAllString(str, `${1} ${2}`)
			str = strings.ToLower(str)
			formattedArgs[i] = str
		} else {
			formattedArgs[i] = arg
		}
	}
	return formattedArgs
}

func (l *ZapLogger) WithError(err error) Logger {
	return &ZapLogger{
		logger: l.logger.With(zap.Error(err)),
	}
}
