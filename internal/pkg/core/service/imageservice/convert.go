package imageservice

import (
	"mini-image-converter/internal/pkg/core/adapter/imageadapter"
	"mini-image-converter/internal/pkg/platform/errors"

	"gocv.io/x/gocv"
)

func (i ImageService) Convert(req imageadapter.ConvertRequest) (imageadapter.ConvertResponse, error) {
	retrieveFileRes, err := i.common.RetrieveFile(req.File, req.FileName, i.os)
	if err != nil {
		return imageadapter.ConvertResponse{}, err
	}
	defer i.os.Remove(retrieveFileRes.FileDestination)

	//Check Content Type
	isPng, err := i.common.IsPNG(i.os, retrieveFileRes.FileDestination)
	if err != nil {
		return imageadapter.ConvertResponse{}, err
	}
	if !isPng {
		return imageadapter.ConvertResponse{}, errors.ErrorNotPng
	}

	readMat := i.gocv.IMRead(retrieveFileRes.FileDestination, gocv.IMReadAnyColor)

	jpegName := retrieveFileRes.FileDestination[:len(retrieveFileRes.FileDestination)-3] + "jpg"

	isSuccess := i.gocv.IMWrite(jpegName, readMat)
	if !isSuccess {
		return imageadapter.ConvertResponse{}, errors.ErrorImWrite
	}
	return imageadapter.ConvertResponse{
		JpegName: jpegName,
	}, nil
}
