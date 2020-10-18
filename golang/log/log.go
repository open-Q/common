package log

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Logger represents geneic logger instance.
type Logger struct {
	*logrus.Logger
}

// NewFileLogger creates new file logger.
func NewFileLogger(logFolder, logFile string, perm os.FileMode) (*Logger, error) {
	if err := os.MkdirAll(logFolder, perm); err != nil && err != os.ErrExist {
		return nil, errors.Wrap(err, "could not create log folder")
	}

	logPath := path.Join(logFolder, logFile)
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open log file %s", logPath)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(f)

	return &Logger{logger}, nil
}
