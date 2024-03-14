package io

import "io"

type IO struct{}

func NewIO() IO {
	return IO{}
}

func (i IO) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}
