package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStorageError_Error(t *testing.T) {
	err := StorageError{
		err: errors.New("some error"),
	}
	require.Error(t, err)
	require.Equal(t, "some error", err.Error())
}

func TestStorageError_Unwrap(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := fmt.Errorf("error 2: %w", err1)
	err := StorageError{
		err: fmt.Errorf("%w: some error", err2),
	}
	require.Error(t, err)
	unwrapErr := errors.Unwrap(err)
	require.Error(t, unwrapErr)
	require.EqualError(t, unwrapErr, "error 2: error 1")
	unwrapErr = errors.Unwrap(unwrapErr)
	require.Error(t, unwrapErr)
	require.EqualError(t, unwrapErr, "error 1")
	unwrapErr = errors.Unwrap(unwrapErr)
	require.NoError(t, unwrapErr)
}

func Test_NewStorageConvertError(t *testing.T) {
	err := NewStorageConvertError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageConvert))
	var e StorageConvertError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageConvert))
}

func Test_NewStorageInsertError(t *testing.T) {
	err := NewStorageInsertError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageInsert))
	var e StorageInsertError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageInsert))
}

func Test_NewStorageDeleteError(t *testing.T) {
	err := NewStorageDeleteError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageDelete))
	var e StorageDeleteError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageDelete))
}

func Test_NewStorageUpdateError(t *testing.T) {
	err := NewStorageUpdateError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageUpdate))
	var e StorageUpdateError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageUpdate))
}

func Test_NewStorageFindError(t *testing.T) {
	err := NewStorageFindError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageFind))
	var e StorageFindError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageFind))
}

func Test_NewStorageUnknownError(t *testing.T) {
	err := NewStorageUnknownError("some error")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrStorageUnknown))
	var e StorageUnknownError
	require.True(t, errors.As(err, &e))
	require.EqualError(t, err, fmt.Sprintf("%v: some error", ErrStorageUnknown))
}
