package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"memory-card-game-BE/pkgs/game"
	"os"
	"sync"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var once sync.Once //nolint:gochecknoglobals, lll, we are creating singleton object, which is going to be reused across the project
var instance *otelzap.Logger

type Logger interface {
	// Debug logs a debug message with the given fields
	Debug(format string, args ...Field)

	// Info logs an info message with the given fields
	Info(format string, args ...Field)

	// Error logs an error message with the given fields
	Error(format string, args ...Field)

	// Warn logs a warning message with the given fields
	Warn(format string, args ...Field)
}

func zapFields(value ...Field) []zap.Field {
	var fields []zap.Field
	for _, f := range value {
		fields = append(fields, f.ZapField())
	}

	return fields
}

// defaultLogger is an abstraction over otelzap logger
// This way internal logger can be replaced easily as required
// otelzap is a thin wrapper for zap.Logger that adds Ctx to logger methods.
// This is useful for tracking trace_id.
// Therefore, logs can be indexed by trace_id in logFn aggregator tools like ELK stack.
// application logger is a singleton, and we can obtain the logger object using `GetLogger` logFn wherever required.
type defaultLogger struct {
	logger otelzap.LoggerWithCtx
}

// GetLogger returns the logger instance.
// If the logger is not initialized, it initializes the logger with info logFn level.
func GetLogger(ctx interface{}) Logger {
	if instance == nil {
		conf := Config{Level: "info", Writer: os.Stdout}
		InitializeLogger(conf)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	if ginCtx, ok := ctx.(*gin.Context); ok {
		if ginCtx == nil {
			return &defaultLogger{logger: instance.Ctx(context.Background())}
		}

		return &defaultLogger{logger: instance.Ctx(ginCtx.Request.Context())}
	}

	return &defaultLogger{logger: instance.Ctx(ctx.(context.Context))}
}

// Debug logs a debug message with the given fields.
func (l *defaultLogger) Debug(msg string, args ...Field) {
	l.logger.Debug(msg, append(zapFields(args...), l.traceDetails()...)...)
}

// Info logs an info message with the given fields.
func (l *defaultLogger) Info(msg string, args ...Field) {
	l.logger.Info(msg, append(zapFields(args...), l.traceDetails()...)...)
}

// Warn logs a warning message with the given fields.
func (l *defaultLogger) Warn(msg string, args ...Field) {
	l.logger.Warn(msg, append(zapFields(args...), l.traceDetails()...)...)
}

// Error logs an error message with the given fields.
func (l *defaultLogger) Error(msg string, args ...Field) {
	l.logger.Error(msg, append(zapFields(args...), l.traceDetails()...)...)
}

func (l *defaultLogger) traceDetails() []zap.Field {
	ctx := l.logger.Context()

	return []zap.Field{traceDetail(ctx, game.CorrelationIdContext), traceDetail(ctx, game.ServiceNameContext)}
}

func traceDetail(ctx context.Context, key game.ContextKey) zap.Field {
	value := fetchDetailsFromContext(ctx, key)

	return zap.String(string(key), value)
}

func fetchDetailsFromContext(ctx context.Context, key any) string {
	val := ctx.Value(key)
	if val != nil {
		if value, ok := val.(string); ok {
			return value
		}
	}

	return ""
}

// newZapLogger returns a new zap logger with the given level.
func newZapLogger(level zapcore.Level, ws ...zapcore.WriteSyncer) *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig()),
		zapcore.NewMultiWriteSyncer(ws...),
		zap.NewAtomicLevelAt(level),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// encoderConfig returns a zapcore.EncoderConfig.
func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
}

// newLogger returns a new logger with the given config.
func newLogger(config Config) *otelzap.Logger {
	return otelzap.New(
		newZapLogger(config.GetLevel(), config.Writer),
		otelzap.WithMinLevel(config.GetLevel()),
		otelzap.WithCallerDepth(1),
	)
}

// InitializeLogger initializes the logger with the given config.
// It is a singleton and can be initialized only once.
func InitializeLogger(config Config) {
	once.Do(func() {
		instance = newLogger(config)
		defer func(zapLogger *otelzap.Logger) { _ = zapLogger.Sync() }(instance)
	})
}
