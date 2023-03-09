// Package logrus provides a BaseLogger implementation for logrus
package logrus

import (
	"context"
	"fmt"
	"github.com/carousell/logging"
	stdlog "log"
	"os"

	log "github.com/sirupsen/logrus"
)

type logger struct {
	logger *log.Logger
	opt    logging.Options
}

func toLogrusLogLevel(level logging.Level) log.Level {
	switch level {
	case logging.DebugLevel:
		return log.DebugLevel
	case logging.InfoLevel:
		return log.InfoLevel
	case logging.WarnLevel:
		return log.WarnLevel
	case logging.ErrorLevel:
		return log.ErrorLevel
	default:
		return log.ErrorLevel
	}
}

func (l *logger) Log(ctx context.Context, level logging.Level, skip int, args ...interface{}) {
	fields := make(log.Fields)

	// fetch fields from context and add them to logrus fields
	ctxFields := logging.FromContext(ctx)
	if ctxFields != nil {
		for k, v := range ctxFields {
			fields[k] = v
		}
	}

	if l.opt.CallerInfo {
		_, file, line := logging.FetchCallerInfo(skip+1, l.opt.CallerFileDepth)
		fields[l.opt.CallerFieldName] = fmt.Sprintf("%s:%d", file, line)
	}

	logger := l.logger.WithFields(fields)
	switch level {
	case logging.DebugLevel:
		logger.Debug(args...)
	case logging.InfoLevel:
		logger.Info(args...)
	case logging.WarnLevel:
		logger.Warn(args...)
	case logging.ErrorLevel:
		logger.Error(args...)
	default:
		l.logger.Error(args...)
	}
}

func (l *logger) SetLevel(level logging.Level) {
	l.logger.SetLevel(toLogrusLogLevel(level))
}

func (l *logger) GetLevel() logging.Level {
	switch l.logger.Level {
	case log.DebugLevel:
		return logging.DebugLevel
	case log.InfoLevel:
		return logging.InfoLevel
	case log.WarnLevel:
		return logging.WarnLevel
	case log.ErrorLevel:
		return logging.ErrorLevel
	default:
		return logging.InfoLevel
	}
}

// NewLogger returns a BaseLogger impl for logrus
func NewLogger(options ...logging.Option) logging.BaseLogger {
	// default options
	opt := logging.GetDefaultOptions()
	// read options
	for _, f := range options {
		f(&opt)
	}

	l := logger{}
	l.logger = log.New()
	l.logger.Out = os.Stdout

	l.logger.SetLevel(toLogrusLogLevel(opt.Level))

	fieldMap := log.FieldMap{
		log.FieldKeyTime:  opt.TimestampFieldName,
		log.FieldKeyLevel: opt.LevelFieldName,
	}
	//check JSON logs
	if opt.JSONLogs {
		l.logger.Formatter = &log.JSONFormatter{
			FieldMap: fieldMap,
		}
	} else {
		l.logger.Formatter = &log.TextFormatter{
			FullTimestamp: true,
		}
	}

	l.opt = opt

	if opt.ReplaceStdLogger {
		stdlog.SetOutput(l.logger.Writer())
	}
	return &l
}
