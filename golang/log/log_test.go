package log

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_NewFileLogger(t *testing.T) {
	tt := []struct {
		name      string
		logFolder string
		logFile   string
		perm      os.FileMode
		before    func(t *testing.T)
		after     func(t *testing.T, logger *Logger)
		expErr    error
	}{
		{
			name:      "create dir error",
			logFolder: "./log/folder",
			perm:      os.ModePerm,
			before: func(t *testing.T) {
				err := os.Mkdir("./log", 0)
				require.NoError(t, err)
			},
			after: func(t *testing.T, logger *Logger) {
				err := os.Chmod("./log", 0777)
				require.NoError(t, err)
				err = os.Remove("./log")
				require.NoError(t, err)
			},
			expErr: errors.Wrap(errors.New("mkdir ./log/folder: permission denied"), "could not create log folder"),
		},
		{
			name:      "open file error",
			logFolder: "./testlog/folder",
			logFile:   "test.json",
			perm:      os.ModePerm,
			before: func(t *testing.T) {
				err := os.MkdirAll("./testlog/folder", os.ModePerm)
				require.NoError(t, err)
				err = os.Chmod("./testlog/folder", 0)
				require.NoError(t, err)
			},
			after: func(t *testing.T, logger *Logger) {
				err := os.Chmod("./testlog/folder", 0777)
				require.NoError(t, err)
				err = os.RemoveAll("./testlog")
				require.NoError(t, err)
			},
			expErr: errors.Wrapf(errors.New("open testlog/folder/test.json: permission denied"), "could not open log file %s", "testlog/folder/test.json"),
		},
		{
			name:      "all ok",
			logFolder: "./testlog/folder",
			logFile:   "test.json",
			perm:      os.ModePerm,
			after: func(t *testing.T, logger *Logger) {
				logger.Infof("hello %s", "test")
				data, err := ioutil.ReadFile("./testlog/folder/test.json")
				require.NoError(t, err)
				require.NotNil(t, data)
				var logMap map[string]interface{}
				err = json.Unmarshal(data, &logMap)
				require.NoError(t, err)
				require.NotNil(t, logMap)
				require.Equal(t, logrus.InfoLevel.String(), logMap["level"])
				require.Equal(t, "hello test", logMap["msg"])
				err = os.RemoveAll("./testlog")
				require.NoError(t, err)
			},
		},
	}
	for i := range tt {
		tc := &tt[i]
		t.Run(tc.name, func(t *testing.T) {
			if tc.before != nil {
				tc.before(t)
			}
			l, err := NewFileLogger(tc.logFolder, tc.logFile, tc.perm)
			if tc.after != nil {
				defer tc.after(t, l)
			}
			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
				return
			}
			require.NotNil(t, l)
		})
	}
}
