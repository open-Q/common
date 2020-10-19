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
	return &cli.BoolFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(bool),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newDurationFlagCli(flag Flag, destination *time.Duration) *cli.DurationFlag {
	return &cli.DurationFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(time.Duration),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newFloat64FlagCli(flag Flag, destination *float64) *cli.Float64Flag {
	return &cli.Float64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(float64),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newInt64FlagCli(flag Flag, destination *int64) *cli.Int64Flag {
	return &cli.Int64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(int64),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newIntFlagCli(flag Flag, destination *int) *cli.IntFlag {
	return &cli.IntFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(int),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newStringFlagCli(flag Flag, destination *string) *cli.StringFlag {
	return &cli.StringFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(string),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newUint64FlagCli(flag Flag, destination *uint64) *cli.Uint64Flag {
	return &cli.Uint64Flag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(uint64),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newUintFlagCli(flag Flag, destination *uint) *cli.UintFlag {
	return &cli.UintFlag{
		Destination: destination,
		Name:        flag.Name,
		Usage:       flag.Usage,
		Value:       flag.Value.(uint),
		Aliases:     flag.Aliases,
		Required:    flag.Required,
	}
}

func newIntSliceFlagCli(flag Flag, destination []int) *cli.IntSliceFlag {
	return &cli.IntSliceFlag{
		Value:    cli.NewIntSlice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
	}
}

func newInt64SliceFlagCli(flag Flag, destination []int64) *cli.Int64SliceFlag {
	return &cli.Int64SliceFlag{
		Value:    cli.NewInt64Slice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
	}
}

func newFloat64SliceFlagCli(flag Flag, destination []float64) *cli.Float64SliceFlag {
	return &cli.Float64SliceFlag{
		Value:    cli.NewFloat64Slice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
	}
}

func newStringSliceFlagCli(flag Flag, destination []string) *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Value:    cli.NewStringSlice(destination...),
		Name:     flag.Name,
		Usage:    flag.Usage,
		Aliases:  flag.Aliases,
		Required: flag.Required,
	}
}
