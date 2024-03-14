package osadapter

import (
	"io"
)

type Adapter interface {
	Remove(path string) error
	ReadFile(path string) ([]byte, error)
	Create(path string) (io.WriteCloser, error)
	Open(path string) (io.ReadCloser, error)
}
