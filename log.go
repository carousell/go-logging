package logging

import (
	"context"
	"sync"
)

var (
	globalLogger Logger
	mu           sync.Mutex
	once         sync.Once
)

type logger struct {
	baseLog BaseLogger
}

func (l *logger) SetLevel(level Level) {
	l.baseLog.SetLevel(level)
}

func (l *logger) GetLevel() Level {
	return l.baseLog.GetLevel()
}

func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	l.Log(ctx, DebugLevel, 1, args...)
}

func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.Log(ctx, InfoLevel, 1, args...)
}

func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	l.Log(ctx, WarnLevel, 1, args...)
}

func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.Log(ctx, ErrorLevel, 1, args...)
}

func (l *logger) Log(ctx context.Context, level Level, skip int, args ...interface{}) {
	if ctx == nil {
		ctx = context.Background()
	}
	if l.GetLevel() >= level {
		l.baseLog.Log(ctx, level, skip+1, args...)
	}
}

// NewLogger creates a new logger with a provided BaseLogger
func NewLogger(log BaseLogger) Logger {
	l := new(logger)
	if log == nil {
		log = NewStdLogger()
	}
	l.baseLog = log
	return l
}

// GetLogger returns the global logger
func GetLogger() Logger {
	if globalLogger == nil {
		mu.Lock()
		defer mu.Unlock()
		globalLogger = NewLogger(nil)
	}
	return globalLogger
}

// SetLogger sets the global logger
func SetLogger(l Logger) {
	if l != nil {
		mu.Lock()
		defer mu.Unlock()
		globalLogger = l
	}
}

// RegisterLogger register a BaseLogger for global use
func RegisterLogger(logger BaseLogger) {
	once.Do(func() {
		if globalLogger == nil {
			globalLogger = NewLogger(logger)
		}
	})
}
