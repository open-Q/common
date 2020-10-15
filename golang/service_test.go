package golang

import (
	"testing"
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/stretchr/testify/require"
)

func Test_generateServiceFlags(t *testing.T) {
	tNow := time.Now()
	tt := []struct {
		name  string
		flags []ServiceFlag
	}{
		{
			name: "bool flags",
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			flags: []ServiceFlag{
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
			name: "all possible flags",
			flags: []ServiceFlag{
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
			cliFlags, flagsMap := generateServiceFlags(tc.flags)
			require.NotNil(t, cliFlags)
			require.NotNil(t, flagsMap)
			service := createTestService([]micro.Option{micro.Flags(cliFlags...)})
			service.Init()
			require.Equal(t, len(tc.flags), len(cliFlags))
			require.Equal(t, len(tc.flags), len(flagsMap))
			for i := range tc.flags {
				checkKey(t, flagsMap, tc.flags[i].Name, tc.flags[i].Value)
			}
		})
	}

	t.Run("invalid flag type", func(t *testing.T) {
		flags := []ServiceFlag{
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

	t.Run("flag redefinition error", func(t *testing.T) {
		defer func() {
			err := recover().(string)
			require.Equal(t, "go.micro.server flag redefined: custom-flag", err)
		}()
		flags := []ServiceFlag{
			{
				Type:  "bool",
				Name:  "custom-flag",
				Value: true,
			},
			{
				Type:  "duration",
				Name:  "custom-flag",
				Value: tNow.Add(40 * time.Millisecond).Sub(tNow),
			},
			{
				Type:  "float64",
				Name:  "custom-flag",
				Value: float64(123.5),
			},
		}
		cliFlags, flagsMap := generateServiceFlags(flags)
		require.NotNil(t, cliFlags)
		require.NotNil(t, flagsMap)
		service := createTestService([]micro.Option{micro.Flags(cliFlags...)})
		service.Init()
	})
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
