package golang

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
)

// ServiceContract represents service contract configuration.
type ServiceContract struct {
	Name        string        `json:"service"`
	Version     string        `json:"version,omitempty"`
	Description string        `json:"description,omitempty"`
	Config      ServiceConfig `json:"config"`
	Flags       []ServiceFlag `json:"flags,omitempty"`
}

// ServiceConfig represents service configuration model.
type ServiceConfig struct {
	Port int64             `json:"port"`
	Host string            `json:"host"`
	Meta map[string]string `json:"meta,omitempty"`
}

// ServiceFlag represents service flag model.
type ServiceFlag struct {
	Type     string      `json:"type,omitempty"`
	Name     string      `json:"name"`
	Value    interface{} `json:"value,omitempty"`
	Usage    string      `json:"usage,omitempty"`
	Aliases  []string    `json:"aliases,omitempty"`
	Required bool        `json:"required,omitempty"`
}

// Validate validates service contract struct.
func (c *ServiceContract) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("service name is required")
	}

	if strings.TrimSpace(c.Config.Host) == "" {
		return errors.New("service host is required")
	}

	for i := range c.Flags {
		if err := c.Flags[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates service flag struct.
func (sf *ServiceFlag) Validate() error {
	if strings.TrimSpace(sf.Name) == "" {
		return errors.New("name is required")
	}

	return nil
}

// NewService creates new micro.Service instance by contract configuration.
func NewService(contractPath string) (micro.Service, map[string]GenericFlag, error) {
	// get service contract.
	contract, err := parseContractFile(contractPath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "contract error")
	}

	// validate contract.
	if err := contract.Validate(); err != nil {
		return nil, nil, errors.Wrap(err, "validation error")
	}

	// create a new service instance.
	cliFlags, flagsMap := generateServiceFlags(contract.Flags)
	service := micro.NewService(
		micro.Name(contract.Name),
		micro.Version(contract.Version),
		micro.Metadata(contract.Config.Meta),
		micro.Flags(cliFlags...),
	)

	// init will parse the command line flags.
	service.Init()

	return service, flagsMap, nil
}

func parseContractFile(fPath string) (*ServiceContract, error) {
	// read the file data.
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read %s file data", fPath)
	}

	// unmarshal file data into the struct.
	var contract ServiceContract
	if err := json.Unmarshal(data, &contract); err != nil {
		return nil, errors.Wrap(err, "could not parse file")
	}

	return &contract, nil
}

func generateServiceFlags(flags []ServiceFlag) ([]cli.Flag, map[string]GenericFlag) {
	cliFlags := make([]cli.Flag, 0, len(flags))
	flagsMap := make(map[string]GenericFlag)

	for i := range flags {
		var dest interface{}
		cliFlag := createFlag(flags[i], &dest)
		if cliFlag == nil {
			continue
		}
		cliFlags = append(cliFlags, cliFlag)
		flagsMap[flags[i].Name] = NewGenericFlag(&dest)
	}

	return cliFlags, flagsMap
}

func createFlag(flag ServiceFlag, destination *interface{}) (cliFlag cli.Flag) {
	switch strings.ToLower(flag.Type) {
	case boolFlag:
		var dest bool
		cliFlag = newBoolFlagCli(flag, &dest)
		*destination = &dest
	case durationFlag:
		var dest time.Duration
		cliFlag = newDurationFlagCli(flag, &dest)
		*destination = &dest
	case float64Flag:
		var dest float64
		cliFlag = newFloat64FlagCli(flag, &dest)
		*destination = &dest
	case int64Flag:
		var dest int64
		cliFlag = newInt64FlagCli(flag, &dest)
		*destination = &dest
	case intFlag:
		var dest int
		cliFlag = newIntFlagCli(flag, &dest)
		*destination = &dest
	case stringFlag:
		var dest string
		cliFlag = newStringFlagCli(flag, &dest)
		*destination = &dest
	case uint64Flag:
		var dest uint64
		cliFlag = newUint64FlagCli(flag, &dest)
		*destination = &dest
	case uintFlag:
		var dest uint
		cliFlag = newUintFlagCli(flag, &dest)
		*destination = &dest
	case intSliceFlag:
		var dest []int
		if flag.Value != nil {
			dest = flag.Value.([]int)
		}
		cliFlag = newIntSliceFlagCli(flag, dest)
		*destination = &dest
	case int64SliceFlag:
		var dest []int64
		if flag.Value != nil {
			dest = flag.Value.([]int64)
		}
		cliFlag = newInt64SliceFlagCli(flag, dest)
		*destination = &dest
	case float64SliceFlag:
		var dest []float64
		if flag.Value != nil {
			dest = flag.Value.([]float64)
		}
		cliFlag = newFloat64SliceFlagCli(flag, dest)
		*destination = &dest
	case stringSliceFlag:
		var dest []string
		if flag.Value != nil {
			dest = flag.Value.([]string)
		}
		cliFlag = newStringSliceFlagCli(flag, dest)
		*destination = &dest
	}

	return
}
