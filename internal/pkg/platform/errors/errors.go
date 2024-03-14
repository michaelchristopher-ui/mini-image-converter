package errors

import "errors"

var (
	ErrorNotPng  = errors.New("uploaded file is not png")
	ErrorImWrite = errors.New("failed on image write")
	ErrFoo       = errors.New("errFoo")
)
