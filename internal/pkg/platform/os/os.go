package os

import (
	"io"
	"os"
)

type OS struct{}

func NewOS() OS {
	return OS{}
}

func (o OS) Remove(path string) error {
	return os.Remove(path)
}

func (o OS) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (o OS) Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

func (o OS) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
