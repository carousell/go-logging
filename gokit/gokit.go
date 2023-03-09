// Package gokit provides BaseLogger implementation for go-kit/log
package gokit

import (
	"context"
	"fmt"
	"github.com/carousell/logging/logging"
	syslog "log"
	"os"

	"github.com/go-kit/kit/log"
)

type logger struct {
	logger log.Logger
	level  logging.Level
	opt    logging.Options
}

func (l *logger) Log(ctx context.Context, level logging.Level, skip int, args ...interface{}) {
	lgr := log.With(l.logger, l.opt.LevelFieldName, level.String())

	if l.opt.CallerInfo {
		_, file, line := logging.FetchCallerInfo(skip+1, l.opt.CallerFileDepth)
		lgr = log.With(lgr, l.opt.CallerFieldName, fmt.Sprintf("%s:%d", file, line))
	}

	// fetch fields from context and add them to logrus fields
	ctxFields := logging.FromContext(ctx)
	if ctxFields != nil {
		for k, v := range ctxFields {
			lgr = log.With(lgr, k, v)
		}
	}

	if len(args) == 1 {
		lgr.Log("msg", args[0])
	} else {
		lgr.Log(args...)
	}
}

func (l *logger) SetLevel(level logging.Level) {
	l.level = level
}

func (l *logger) GetLevel() logging.Level {
	return l.level
}

// NewLogger returns a base logger impl for go-kit log
func NewLogger(options ...logging.Option) logging.BaseLogger {
	// default options
	opt := logging.GetDefaultOptions()

	// read options
	for _, f := range options {
		f(&opt)
	}

	l := logger{}
	writer := log.NewSyncWriter(os.Stdout)

	// check for json or logfmt
	if opt.JSONLogs {
		l.logger = log.NewJSONLogger(writer)
	} else {
		l.logger = log.NewLogfmtLogger(writer)
	}

	l.logger = log.With(l.logger, opt.TimestampFieldName, log.DefaultTimestamp)

	l.level = opt.Level
	l.opt = opt

	if opt.ReplaceStdLogger {
		syslog.SetFlags(syslog.LUTC)
		syslog.SetOutput(log.NewStdlibAdapter(l.logger, log.TimestampKey(opt.TimestampFieldName)))
	}
	return &l
}
