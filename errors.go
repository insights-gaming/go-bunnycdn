package bunnycdn

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound           = errors.New("bunny: not found")
	ErrUploadFailed       = errors.New("bunny: upload failed")
	ErrObjectDeleteFailed = errors.New("bunny: object delete failed")
	ErrInvalidURI         = errors.New("bunny: invalid uri")
)

type Error struct {
	HttpCode int
	Message  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("bunny: %d %s", e.HttpCode, e.Message)
}

type UnexpectedResponseError struct {
	Code   int
	Status string
	Body   []byte
}

func (e *UnexpectedResponseError) Error() string {
	return fmt.Sprintf("bunny: unexpected response - %d %s\n\n%q", e.Code, e.Status, string(e.Body))
}

