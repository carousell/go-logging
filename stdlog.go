// Package stdlog provides a BaseLogger implementation for golang "log" package
package logging

import (
	"context"
	"log"
)

type stdlogger struct {
	level Level
}

func (l *stdlogger) SetLevel(level Level) {
	l.level = level
}

func (l *stdlogger) GetLevel() Level {
	return l.level
}

func (l *stdlogger) Log(ctx context.Context, level Level, skip int, args ...interface{}) {
	if l.level >= level {
		ctxFields := FromContext(ctx)
		if ctxFields != nil {
			for k, v := range ctxFields {
				args = append(args, k, v)
			}
		}
		log.Println(args...)
	}
}

// NewStdLogger returns a BaseLogger impl for golang "log" package
func NewStdLogger(options ...Option) BaseLogger {
	return &stdlogger{
		level: InfoLevel,
	}
}
