package logging

import (
	"context"
)

// Logger interface is implemented by the log implementation
type Logger interface {
	BaseLogger
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
}
