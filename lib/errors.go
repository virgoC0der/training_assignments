package lib

import "errors"

var (
	ErrIDExist         = errors.New("id already exists")
	ErrIDNotFound      = errors.New("id not found")
	ErrNameNotFound    = errors.New("name not found")
	ErrIllegalInput    = errors.New("illegal input")
	ErrHandlerNotFound = errors.New("handler not found")
)
