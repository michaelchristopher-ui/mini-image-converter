package common

import (
	"mini-image-converter/internal/pkg/core/adapter/commonadapter"
	"mini-image-converter/internal/pkg/core/adapter/ioadapter"
	"net/http"
)

type Common struct {
	io ioadapter.Adapter
}

func NewCommon(io ioadapter.Adapter) Common {
	return Common{
		io: io,
	}
}

func (c Common) RetrieveFile(req commonadapter.FileOpener, filename string, os commonadapter.FileCreator) (commonadapter.RetrieveFileRes, error) {
	src, err := req.Open()
	if err != nil {
		return commonadapter.RetrieveFileRes{}, err
	}
	defer src.Close()

	dst, err := os.Create(filename)
	if err != nil {
		return commonadapter.RetrieveFileRes{}, err
	}
	defer dst.Close()

	if _, err = c.io.Copy(dst, src); err != nil {
		return commonadapter.RetrieveFileRes{}, err
	}
	return commonadapter.RetrieveFileRes{
		FileDestination: filename,
	}, nil
}

func (c Common) IsPNG(os commonadapter.FileReader, fileString string) (bool, error) {
	data, err := os.ReadFile(fileString)
	if err != nil {
		return false, err
	}
	contentString := http.DetectContentType(data)
	if contentString != "image/png" {
		return false, nil
	}
	return true, nil
}
