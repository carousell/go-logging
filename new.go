package logging

import (
	"context"
)

type Attr [2]interface{}

type AttrMap map[string]interface{}

type LoggerV2 interface {
	Info(ctx context.Context, msg string, attrs ...Attr)
}

func NewLoggerV2(baseLogger BaseLogger) LoggerV2 {
	return loggerV2{
		baseLogger: baseLogger,
	}
}

type loggerV2 struct {
	baseLogger BaseLogger

	attrs []Attr
}

func (l loggerV2) log(ctx context.Context, level Level, msg string, attrs ...Attr) {
	args := attrsToArgs(l.attrs)
	args = append(args, attrsToArgs(attrs)...)
	args = append(args, "msg", msg)
	l.baseLogger.Log(ctx, level, 1, args...)
}

func (l loggerV2) Info(ctx context.Context, msg string, attrs ...Attr) {
	l.log(ctx, InfoLevel, msg, attrs...)
}

type loggerKeyType struct{}

var loggerKey loggerKeyType = struct{}{}

func WithLogger(ctx context.Context, attrs ...Attr) context.Context {
	if parentLogger, ok := GetLoggerV2(ctx).(loggerV2); ok {
		attrs = append(attrs, parentLogger.attrs...)
	}

	logger := loggerV2{
		baseLogger: globalLogger,
		attrs:      attrs,
	}

	ctx = context.WithValue(ctx, loggerKey, logger)

	return ctx
}

func GetLoggerV2(ctx context.Context) LoggerV2 {
	logger, ok := ctx.Value(loggerKey).(LoggerV2)
	if !ok {
		return NewLoggerV2(globalLogger)
	}
	return logger
}

func InfoV2(ctx context.Context, msg string, attrs ...Attr) {
	logger := GetLoggerV2(ctx)
	logger.Info(ctx, msg, attrs...)
}

func InfoV3(ctx context.Context, msg string, attrs AttrMap) {
	logger := GetLoggerV2(ctx)
	var args []Attr
	for k, v := range attrs {
		args = append(args, WithAttr(k, v))
	}
	logger.Info(ctx, msg, args...)
}

func WithAttr(key string, value interface{}) Attr {
	return [2]interface{}{key, value}
}

func attrsToArgs(attrs []Attr) []interface{} {
	args := make([]interface{}, 0, len(attrs)*2)
	for _, attr := range attrs {
		args = append(args, attr[0], attr[1])
	}
	return args
}
