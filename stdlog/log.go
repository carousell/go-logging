// Package stdlog provides a BaseLogger implementation for golang "log" package
package stdlog

import (
	"context"
	"github.com/carousell/logging/logging"
	"log"
)

type logger struct {
	level logging.Level
}

func (l *logger) SetLevel(level logging.Level) {
	l.level = level
}

func (l *logger) GetLevel() logging.Level {
	return l.level
}

func (l *logger) Log(ctx context.Context, level logging.Level, skip int, args ...interface{}) {
	if l.level >= level {
		// fetch fields from context and add them to logrus fields
		ctxFields := logging.FromContext(ctx)
		if ctxFields != nil {
			for k, v := range ctxFields {
				args = append(args, k, v)
			}
		}
		log.Println(args...)
	}
}

// NewLogger returns a BaseLogger impl for golang "log" package
func NewLogger(options ...logging.Option) logging.BaseLogger {
	return &logger{
		level: logging.InfoLevel,
	}
}
