package logging

import (
	"context"
	"runtime"
	"strings"
)

// BaseLogger is the interface that needs to be implemented by client loggers
type BaseLogger interface {
	Log(ctx context.Context, level Level, skip int, args ...interface{})
	SetLevel(level Level)
	GetLevel() Level
}

// Options contain all common options for BaseLoggers
type Options struct {
	ReplaceStdLogger   bool
	JSONLogs           bool
	Level              Level
	TimestampFieldName string
	LevelFieldName     string
	CallerInfo         bool
	CallerFileDepth    int
	CallerFieldName    string
}

// GetDefaultOptions fetches loggers default options
func GetDefaultOptions() Options {
	return DefaultOptions
}

// DefaultOptions stores all default options in loggers package
var (
	DefaultOptions = Options{
		ReplaceStdLogger:   false,
		JSONLogs:           true,
		Level:              InfoLevel,
		TimestampFieldName: "@timestamp",
		LevelFieldName:     "level",
		CallerInfo:         true,
		CallerFileDepth:    2,
		CallerFieldName:    "caller",
	}
)

// Option defines an option for BaseLogger
type Option func(*Options)

// WithReplaceStdLogger enables/disables replacing std logger
func WithReplaceStdLogger(replaceStdLogger bool) Option {
	return func(o *Options) {
		o.ReplaceStdLogger = replaceStdLogger
	}
}

// WithJSONLogs enables/disables json logs
func WithJSONLogs(json bool) Option {
	return func(o *Options) {
		o.JSONLogs = json
	}
}

// WithTimestampFieldName sets the name of the time stamp field in logs
func WithTimestampFieldName(name string) Option {
	return func(o *Options) {
		if name != "" {
			o.TimestampFieldName = name
		}
	}
}

// WithLevelFieldName sets the name of the level field in logs
func WithLevelFieldName(name string) Option {
	return func(o *Options) {
		if name != "" {
			o.LevelFieldName = name
		}
	}
}

// WithCallerInfo enables/disables adding caller info to logs
func WithCallerInfo(callerInfo bool) Option {
	return func(o *Options) {
		o.CallerInfo = callerInfo
	}
}

// WithCallerFileDepth sets the depth of file to use in caller info
func WithCallerFileDepth(depth int) Option {
	return func(o *Options) {
		if depth > 0 {
			o.CallerFileDepth = depth
		}
	}
}

// WithCallerFieldName sets the name of callerinfo field
func WithCallerFieldName(name string) Option {
	return func(o *Options) {
		name = strings.TrimSpace(name)
		if name != "" {
			o.CallerFieldName = name
		}
	}
}

// FetchCallerInfo fetches function name, file name and line number from stack
func FetchCallerInfo(skip int, depth int) (function string, file string, line int) {
	if depth <= 0 {
		depth = 2
	}
	if skip < 0 {
		skip = 0
	}
	pc, file, line, _ := runtime.Caller(skip + 1)
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			depth--
			if depth == 0 {
				file = file[i+1:]
				break
			}
		}
	}
	return runtime.FuncForPC(pc).Name(), file, line
}
