package commonadapter

import (
	"io"
	"mime/multipart"
)

type Adapter interface {
	RetrieveFile(req FileOpener, filename string, creator FileCreator) (RetrieveFileRes, error)
	IsPNG(fileReader FileReader, fileString string) (bool, error)
}

type RetrieveFileRes struct {
	FileDestination string
}

type FileOpener interface {
	Open() (multipart.File, error)
}

type FileCreator interface {
	Create(dst string) (io.WriteCloser, error)
}

type FileReader interface {
	ReadFile(string) ([]byte, error)
}
