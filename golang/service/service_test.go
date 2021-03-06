package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_generateServiceFlags(t *testing.T) {
	tNow := time.Now()
	tt := []struct {
		name     string
		flags    []Flag
		prepare  func()
		expValue func(flag *Flag) interface{}
	}{
		{
			name: "bool flags",
			flags: []Flag{
				{
					Type:  "bool",
					Name:  "bool-flag-with-true",
					Value: true,
				},
				{
					Type:  "bool",
					Name:  "bool-flag-with-false",
					Value: false,
				},
			},
		},
		{
			name: "duration flags",
			flags: []Flag{
				{
					Type:  "duration",
					Name:  "duration-flag-with-25ms",
					Value: tNow.Add(25 * time.Millisecond).Sub(tNow),
				},
				{
					Type:  "duration",
					Name:  "duration-flag-with-1h23m43s8ms",
					Value: tNow.Add(time.Hour + 23*time.Minute + 43*time.Second + 8*time.Millisecond).Sub(tNow),
				},
			},
		},
		{
			name: "float64 flags",
			flags: []Flag{
				{
					Type:  "float64",
					Name:  "float64-flag-with-positive",
					Value: float64(123.566),
				},
				{
					Type:  "float64",
					Name:  "float64-flag-with-negative",
					Value: float64(-2222.89),
				},
			},
		},
		{
			name: "int64 flags",
			flags: []Flag{
				{
					Type:  "int64",
					Name:  "int64-flag-with-positive",
					Value: int64(123),
				},
				{
					Type:  "int64",
					Name:  "int64-flag-with-negative",
					Value: int64(-2222),
				},
			},
		},
		{
			name: "int flags",
			flags: []Flag{
				{
					Type:  "int",
					Name:  "int-flag-with-positive",
					Value: int(123),
				},
				{
					Type:  "int",
					Name:  "int-flag-with-negative",
					Value: int(-2222),
				},
			},
		},
		{
			name: "string flags",
			flags: []Flag{
				{
					Type:  "string",
					Name:  "string-flag-with-empty",
					Value: "",
				},
				{
					Type:  "string",
					Name:  "string-flag-with-non-empty",
					Value: "hello\nworld",
				},
			},
		},
		{
			name: "uint64 flags",
			flags: []Flag{
				{
					Type:  "uint64",
					Name:  "uint64-flag-with-zero",
					Value: uint64(0),
				},
				{
					Type:  "uint64",
					Name:  "uint64-flag-with-non-zero",
					Value: uint64(123),
				},
			},
		},
		{
			name: "uint flags",
			flags: []Flag{
				{
					Type:  "uint",
					Name:  "uint-flag-with-zero",
					Value: uint(0),
				},
				{
					Type:  "uint",
					Name:  "uint-flag-with-non-zero",
					Value: uint(123),
				},
			},
		},
		{
			name: "int slice flags",
			flags: []Flag{
				{
					Type:  "slice:int",
					Name:  "slice:int-flag-with-empty",
					Value: []int{},
				},
				{
					Type:  "slice:int",
					Name:  "slice:int-flag-with-non-empty",
					Value: []int{-1, 23, 8, 0, 123},
				},
			},
		},
		{
			name: "int64 slice flags",
			flags: []Flag{
				{
					Type:  "slice:int64",
					Name:  "slice:int64-flag-with-empty",
					Value: []int64{},
				},
				{
					Type:  "slice:int64",
					Name:  "slice:int64-flag-with-non-empty",
					Value: []int64{-1, 23, 8, 0, 123},
				},
			},
		},
		{
			name: "float64 slice flags",
			flags: []Flag{
				{
					Type:  "slice:float64",
					Name:  "slice:float64-flag-with-empty",
					Value: []float64{},
				},
				{
					Type:  "slice:float64",
					Name:  "slice:float64-flag-with-non-empty",
					Value: []float64{-1.3, .23, 8, 40, 12.123},
				},
			},
		},
		{
			name: "string slice flags",
			flags: []Flag{
				{
					Type:  "slice:string",
					Name:  "slice:string-flag-with-empty",
					Value: []string{},
				},
				{
					Type:  "slice:string",
					Name:  "slice:string-flag-with-non-empty",
					Value: []string{"hello", "world"},
				},
			},
		},
		{
			name: "flags form env",
			flags: []Flag{
				{
					Type:         "string",
					Name:         "string-env",
					EnvVariables: []string{"env1"},
				},
				{
					Type:         "float64",
					Name:         "float64-env",
					EnvVariables: []string{"env2"},
				},
			},
			prepare: func() {
				err := os.Setenv("env1", "hello")
				require.NoError(t, err)
				err = os.Setenv("env2", "325.677")
				require.NoError(t, err)
			},
			expValue: func(f *Flag) interface{} {
				if f.Name == "string-env" {
					return "hello"
				}
				return float64(325.677)
			},
		},
		{
			name: "all possible flags",
			flags: []Flag{
				{
					Type:  "bool",
					Name:  "bool-flag",
					Value: true,
				},
				{
					Type:  "duration",
					Name:  "duration-flag",
					Value: tNow.Add(40 * time.Millisecond).Sub(tNow),
				},
				{
					Type:  "float64",
					Name:  "float64-flag",
					Value: float64(123.5),
				},
				{
					Type:  "int64",
					Name:  "int64-flag",
					Value: int64(123),
				},
				{
					Type:  "int",
					Name:  "int-flag",
					Value: int(123),
				},
				{
					Type:  "string",
					Name:  "string-flag",
					Value: "hello",
				},
				{
					Type:  "uint64",
					Name:  "uint64-flag",
					Value: uint64(123),
				},
				{
					Type:  "uint",
					Name:  "uint-flag",
					Value: uint(123),
				},
				{
					Type:  "slice:int",
					Name:  "slice:int-flag",
					Value: []int{1, 2, 3},
				},
				{
					Type:  "slice:int64",
					Name:  "slice:int64-flag",
					Value: []int64{1, 2, 3},
				},
				{
					Type:  "slice:float64",
					Name:  "slice:float64-flag",
					Value: []float64{1.3, -2.0, .3},
				},
				{
					Type:  "slice:string",
					Name:  "slice:string-flag",
					Value: []string{"1", "2", "3"},
				},
			},
		},
	}
	for i := range tt {
		tc := &tt[i]
		t.Run(tc.name, func(t *testing.T) {
			if tc.prepare != nil {
				tc.prepare()
			}
			cliFlags, flagsMap := generateServiceFlags(tc.flags)
			require.NotNil(t, cliFlags)
			require.NotNil(t, flagsMap)
			service := createTestService([]micro.Option{micro.Flags(cliFlags...)})
			service.Init()
			require.Equal(t, len(tc.flags), len(cliFlags))
			require.Equal(t, len(tc.flags), len(flagsMap))
			for i := range tc.flags {
				val := tc.flags[i].Value
				if tc.expValue != nil {
					val = tc.expValue(&tc.flags[i])
				}
				checkKey(t, flagsMap, tc.flags[i].Name, val)
			}
		})
	}

	t.Run("invalid flag type", func(t *testing.T) {
		flags := []Flag{
			{
				Type:  "invalid",
				Name:  "invalid-flag",
				Value: true,
			},
			{
				Type:  "float64",
				Name:  "custom-float64-flag",
				Value: float64(123.5),
			},
		}
		cliFlags, flagsMap := generateServiceFlags(flags)
		require.NotNil(t, cliFlags)
		require.NotNil(t, flagsMap)
		service := createTestService([]micro.Option{micro.Flags(cliFlags...)})
		service.Init()
		require.Equal(t, 1, len(cliFlags))
		require.Equal(t, 1, len(flagsMap))
		checkKey(t, flagsMap, "custom-float64-flag", float64(123.5))
	})
}

func TestContract_Validate(t *testing.T) {
	tt := []struct {
		name     string
		contract Contract
		expErr   error
	}{
		{
			name:     "service name is required error",
			contract: Contract{},
			expErr:   errors.New("service name is required"),
		},
		{
			name: "service host is required error",
			contract: Contract{
				Name: "test",
			},
			expErr: errors.New("service host is required"),
		},
		{
			name: "flags validation error",
			contract: Contract{
				Name: "test",
				Config: Config{
					Host: "127.0.0.1",
				},
				Flags: []Flag{{}},
			},
			expErr: errors.New("flag's name is required"),
		},
		{
			name: "all ok",
			contract: Contract{
				Name: "test",
				Config: Config{
					Host: "127.0.0.1",
				},
				Flags: []Flag{
					{
						Name: "test-flag",
					},
				},
			},
		},
	}
	for i := range tt {
		tc := &tt[i]
		t.Run(tc.name, func(t *testing.T) {
			err := tc.contract.Validate()
			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestFlag_Validate(t *testing.T) {
	t.Run("flag's name is required error", func(t *testing.T) {
		f := Flag{}
		err := f.Validate()
		require.Error(t, err)
		require.EqualError(t, err, errors.New("flag's name is required").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		f := Flag{
			Name: "test-flag",
		}
		err := f.Validate()
		require.NoError(t, err)
	})
}

func Test_parseContractFile(t *testing.T) {
	t.Run("read file error", func(t *testing.T) {
		_, err := parseContractFile("nonexists")
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrapf(errors.New("open nonexists: no such file or directory"), "could not read %s file data", "nonexists").Error())
	})
	t.Run("parse file data error", func(t *testing.T) {
		fPath := path.Join(os.TempDir(), "temp.json")
		err := createFileWithContent(fPath, []byte("invalid data"))
		require.NoError(t, err)
		defer func() {
			err := os.Remove(fPath)
			require.NoError(t, err)
		}()
		_, err = parseContractFile(fPath)
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("invalid character 'i' looking for beginning of value"), "could not parse file").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		contract := Contract{
			Description: "some description",
			Config: Config{
				Port: 9000,
				Host: "127.0.0.1",
				Meta: map[string]string{
					"key": "value",
				},
			},
			Flags: []Flag{
				{
					Name:     "flag-name",
					Aliases:  []string{"fname", "fn"},
					Required: true,
					Type:     "string",
					Usage:    "use for testing",
					Value:    "test value",
				},
			},
			Name:    "test",
			Version: "0.0.1",
		}
		data, err := json.Marshal(contract)
		require.NoError(t, err)
		fPath := path.Join(os.TempDir(), "temp.json")
		err = createFileWithContent(fPath, data)
		require.NoError(t, err)
		defer func() {
			err := os.Remove(fPath)
			require.NoError(t, err)
		}()
		ctr, err := parseContractFile(fPath)
		require.NoError(t, err)
		require.Equal(t, contract, *ctr)
	})
}

func Test_New(t *testing.T) {
	t.Run("parse file error", func(t *testing.T) {
		_, _, err := New("nonexists")
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("could not read nonexists file data: open nonexists: no such file or directory"), "contract error").Error())
	})
	t.Run("validate contract error", func(t *testing.T) {
		contract := Contract{}
		data, err := json.Marshal(contract)
		require.NoError(t, err)
		fPath := path.Join(os.TempDir(), "temp.json")
		err = createFileWithContent(fPath, data)
		require.NoError(t, err)
		defer func() {
			err := os.Remove(fPath)
			require.NoError(t, err)
		}()
		_, _, err = New(fPath)
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("service name is required"), "validation error").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		contract := Contract{
			Description: "some description",
			Config: Config{
				Port: 9000,
				Host: "127.0.0.1",
				Meta: map[string]string{
					"key": "value",
				},
			},
			Flags: []Flag{
				{
					Name:     "flag-name",
					Aliases:  []string{"fname", "fn"},
					Required: true,
					Type:     "string",
					Usage:    "use for testing",
					Value:    "test value",
				},
			},
			Name:    "test",
			Version: "0.0.1",
		}
		data, err := json.Marshal(contract)
		require.NoError(t, err)
		fPath := path.Join(os.TempDir(), "temp.json")
		err = createFileWithContent(fPath, data)
		require.NoError(t, err)
		defer func() {
			err := os.Remove(fPath)
			require.NoError(t, err)
		}()
		err = os.Setenv(disableFlagCheckENV, "true")
		require.NoError(t, err)
		service, flagsMap, err := New(fPath)
		require.NoError(t, err)
		require.NotNil(t, service)
		require.NotNil(t, flagsMap)
		checkKey(t, flagsMap, "flag-name", "test value")
	})
}

func createFileWithContent(fName string, data []byte) error {
	return ioutil.WriteFile(fName, data, 0666)
}

func createTestService(opts []micro.Option) micro.Service {
	service := micro.NewService(opts...)

	service.Options().Cmd.App().OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
		// skip flag parse errors.
		return nil
	}

	return service
}

func checkKey(t *testing.T, m map[string]GenericFlag, key string, expValue interface{}) {
	v, ok := m[key]
	require.True(t, ok)
	if expValue != nil {
		require.NotNil(t, v)
		require.Equal(t, expValue, v.Value())
	}
}
