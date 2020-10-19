package service

import (
	"reflect"
	"time"

	"github.com/micro/cli/v2"
)

const (
	boolFlag         = "bool"
	durationFlag     = "duration"
	float64Flag      = "float64"
	int64Flag        = "int64"
	intFlag          = "int"
	stringFlag       = "string"
	uint64Flag       = "uint64"
	uintFlag         = "uint"
	intSliceFlag     = "slice:int"
	int64SliceFlag   = "slice:int64"
	float64SliceFlag = "slice:float64"
	stringSliceFlag  = "slice:string"
)

// GenericFlag represents generic flag model.
type GenericFlag struct {
	v *interface{}
}

// Value returns actual generic flag value.
// Uses reflection to determinate actual value.
func (gf *GenericFlag) Value() interface{} {
	rv := reflect.ValueOf(gf.v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	if rv.IsValid() {
		return rv.Interface()
	}
	return nil
}

// NewGenericFlag returns new generic flag instance with the provided value.
func NewGenericFlag(value interface{}) GenericFlag {
	return GenericFlag{
		v: &value,
	}
}

func newBoolFlagCli(flag Flag, destination *bool) *cli.BoolFlag {
	f := cli.BoolFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(bool)
	}
	return &f
}

func newDurationFlagCli(flag Flag, destination *time.Duration) *cli.DurationFlag {
	f := cli.DurationFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(time.Duration)
	}
	return &f
}

func newFloat64FlagCli(flag Flag, destination *float64) *cli.Float64Flag {
	f := cli.Float64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(float64)
	}
	return &f
}

func newInt64FlagCli(flag Flag, destination *int64) *cli.Int64Flag {
	f := cli.Int64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(int64)
	}
	return &f
}

func newIntFlagCli(flag Flag, destination *int) *cli.IntFlag {
	f := cli.IntFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(int)
	}
	return &f
}

func newStringFlagCli(flag Flag, destination *string) *cli.StringFlag {
	f := cli.StringFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(string)
	}
	return &f
}

func newUint64FlagCli(flag Flag, destination *uint64) *cli.Uint64Flag {
	f := cli.Uint64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(uint64)
	}
	return &f
}

func newUintFlagCli(flag Flag, destination *uint) *cli.UintFlag {
	f := cli.UintFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Aliases:     flag.Aliases,
		Required:    flag.Required,
		EnvVars:     flag.EnvVariables,
	}
	if flag.Value != nil {
		f.Value = flag.Value.(uint)
	}
	return &f
}

func newIntSliceFlagCli(flag Flag, destination []int) *cli.IntSliceFlag {
	return &cli.IntSliceFlag{
		Value:    cli.NewIntSlice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
		EnvVars:  flag.EnvVariables,
	}
}

func newInt64SliceFlagCli(flag Flag, destination []int64) *cli.Int64SliceFlag {
	return &cli.Int64SliceFlag{
		Value:    cli.NewInt64Slice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
		EnvVars:  flag.EnvVariables,
	}
}

func newFloat64SliceFlagCli(flag Flag, destination []float64) *cli.Float64SliceFlag {
	return &cli.Float64SliceFlag{
		Value:    cli.NewFloat64Slice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
		EnvVars:  flag.EnvVariables,
	}
}

func newStringSliceFlagCli(flag Flag, destination []string) *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Value:    cli.NewStringSlice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
		EnvVars:  flag.EnvVariables,
	}
}
