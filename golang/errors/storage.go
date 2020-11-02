package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// There are common storage errors.
var (
	ErrStorageConvert = errors.New("could not convert value")
	ErrStorageInsert  = errors.New("could not insert value")
	ErrStorageDelete  = errors.New("could not delete value")
	ErrStorageUpdate  = errors.New("could not update value")
	ErrStorageFind    = errors.New("could not find value")
	ErrStorageUnknown = errors.New("unknown error")
)

// StorageError represents common storage error structure.
type StorageError struct {
	err error
}

type (
	// StorageConvertError represents convert value error.
	StorageConvertError struct {
		*StorageError
	}
	// StorageInsertError represents insert value error.
	StorageInsertError struct {
		*StorageError
	}
	// StorageDeleteError represents delete value error.
	StorageDeleteError struct {
		*StorageError
	}
	// StorageUpdateError represents update value error.
	StorageUpdateError struct {
		*StorageError
	}
	// StorageFindError represents find value error.
	StorageFindError struct {
		*StorageError
	}
	// StorageUnknownError represents unknown storage error.
	StorageUnknownError struct {
		*StorageError
	}
)

// Error returns error as a string value.
func (e StorageError) Error() string {
	return e.err.Error()
}

// Unwrap returns the low level of the provided error.
func (e StorageError) Unwrap() error {
	return errors.Unwrap(e.err)
}

// NewStorageConvertError creates new StorageConvertError instance.
func NewStorageConvertError(message string) StorageConvertError {
	return StorageConvertError{newStorageError(message, ErrStorageConvert)}
}

// NewStorageInsertError creates new StorageInsertError instance.
func NewStorageInsertError(message string) StorageInsertError {
	return StorageInsertError{newStorageError(message, ErrStorageInsert)}
}

// NewStorageDeleteError creates new StorageDeleteError instance.
func NewStorageDeleteError(message string) StorageDeleteError {
	return StorageDeleteError{newStorageError(message, ErrStorageDelete)}
}

// NewStorageUpdateError creates new StorageUpdateError instance.
func NewStorageUpdateError(message string) StorageUpdateError {
	return StorageUpdateError{newStorageError(message, ErrStorageUpdate)}
}

// NewStorageFindError creates new StorageFindError instance.
func NewStorageFindError(message string) StorageFindError {
	return StorageFindError{newStorageError(message, ErrStorageFind)}
}

// NewStorageUnknownError creates new StorageUnknownError instance.
func NewStorageUnknownError(message string) StorageUnknownError {
	return StorageUnknownError{newStorageError(message, ErrStorageUnknown)}
}

func newStorageError(msg string, err error) *StorageError {
	return &StorageError{
		err: fmt.Errorf("%w: %s", err, msg),
	}
}
