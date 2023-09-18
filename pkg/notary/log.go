package notary

import (
	"fmt"

	"github.com/go-logr/logr"
	notationlog "github.com/notaryproject/notation-go/log"
)

func NotaryLoggerAdapter(logger logr.Logger) notationlog.Logger {
	return &notaryLoggerAdapter{
		logger: logger.V(4),
	}
}

type notaryLoggerAdapter struct {
	logger logr.Logger
}

func (nla *notaryLoggerAdapter) Debug(args ...interface{}) {
	nla.info(args...)
}

func (nla *notaryLoggerAdapter) Debugf(format string, args ...interface{}) {
	nla.infof(format, args...)
}

func (nla *notaryLoggerAdapter) Debugln(args ...interface{}) {
	nla.infoln(args...)
}

func (nla *notaryLoggerAdapter) Info(args ...interface{}) {
	nla.info(args...)
}

func (nla *notaryLoggerAdapter) Infof(format string, args ...interface{}) {
	nla.infof(format, args...)
}

func (nla *notaryLoggerAdapter) Infoln(args ...interface{}) {
	nla.infoln(args...)
}

func (nla *notaryLoggerAdapter) Warn(args ...interface{}) {
	nla.info(args...)
}

func (nla *notaryLoggerAdapter) Warnf(format string, args ...interface{}) {
	nla.infof(format, args...)
}

func (nla *notaryLoggerAdapter) Warnln(args ...interface{}) {
	nla.infoln(args...)
}

func (nla *notaryLoggerAdapter) Error(args ...interface{}) {
	nla.logger.Error(fmt.Errorf(fmt.Sprint(args...)), "")
}

func (nla *notaryLoggerAdapter) Errorf(format string, args ...interface{}) {
	nla.logger.Error(fmt.Errorf(format, args...), "")
}

func (nla *notaryLoggerAdapter) Errorln(args ...interface{}) {
	nla.logger.Error(fmt.Errorf(fmt.Sprintln(args...)), "")
}

func (nla *notaryLoggerAdapter) info(level int, args ...interface{}) {
	nla.logger.V(level).Info(fmt.Sprint(args...))
}

func (nla *notaryLoggerAdapter) infof(level int, format string, args ...interface{}) {
	nla.logger.V(level).Info(fmt.Sprintf(format, args...))
}

func (nla *notaryLoggerAdapter) infoln(level int, args ...interface{}) {
	nla.logger.V(level).Info(fmt.Sprintln(args...))
}
