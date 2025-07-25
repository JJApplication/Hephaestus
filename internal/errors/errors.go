package errors

import "errors"

var (
	ErrFileSize    = errors.New("file size is too big")
	ErrFileInvalid = errors.New("file is invalid")
	ErrFileNoEqual = errors.New("file has no equal with pkg config")

	ErrEmptyDeploy = errors.New("empty deploy path")
	ErrEmptyApp    = errors.New("empty app name")
)
