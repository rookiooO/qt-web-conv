package log

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type Option struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`
	LogInConsole bool   `yaml:"log-in-console"`
	MaxAge       int64  `yaml:"max-age"`
	CutTime      int64  `yaml:"cut-time"`
}

func New(opt *Option) (*logrus.Entry, error) {
	var (
		logger *logrus.Logger
		level  logrus.Level
		err    error
	)
	logger = logrus.New()
	if opt.Director != "" {
		r, _ := rotatelogs.New(
			opt.Director+".%Y%m%d",
			rotatelogs.WithLinkName(opt.Director),
			rotatelogs.WithMaxAge(time.Hour*24*time.Duration(opt.MaxAge)),
			rotatelogs.WithRotationTime(time.Hour*24*time.Duration(opt.CutTime)),
		)
		logger.SetOutput(r)
	}
	logger.SetReportCaller(opt.ShowLine)
	logger.SetFormatter(opt)
	level, err = logrus.ParseLevel(opt.Level)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}
	return logrus.NewEntry(logger), err
}

func (l *Option) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		color     int
		buf       *bytes.Buffer
		timestamp string
		err       error
		caller    string
	)

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		color = gray
	case logrus.WarnLevel:
		color = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = red
	default:
		color = blue
	}

	timestamp = time.Now().Format("2006-1-2 15:04:05")

	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}

	if entry.HasCaller() {
		caller = fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	}

	if l.Director != "" {
		_, err = fmt.Fprintf(buf, "[%s][%s][%s]-line:%s-message:%s\n",
			l.Prefix, timestamp, entry.Level, caller, entry.Message)
		if err != nil {
			return nil, err
		}
	}
	if l.LogInConsole {
		_, err = fmt.Printf("\u001B[%dm[%s][%s][%s]\nline: %s\nmessage: %s\u001B[0m\n",
			color, l.Prefix, entry.Level, timestamp, caller, entry.Message)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
