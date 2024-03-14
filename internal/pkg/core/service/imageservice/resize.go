package imageservice

import (
	"image"
	"mini-image-converter/internal/pkg/core/adapter/imageadapter"
	"mini-image-converter/internal/pkg/platform/errors"

	"gocv.io/x/gocv"
)

func (i ImageService) Resize(req imageadapter.ResizeRequest) (imageadapter.ResizeResponse, error) {
	retrieveFileRes, err := i.common.RetrieveFile(req.File, req.FileName, i.os)
	if err != nil {
		return imageadapter.ResizeResponse{}, err
	}
	defer i.os.Remove(retrieveFileRes.FileDestination)

	//Check Content Type
	isPng, err := i.common.IsPNG(i.os, retrieveFileRes.FileDestination)
	if err != nil {
		return imageadapter.ResizeResponse{}, err
	}
	if !isPng {
		return imageadapter.ResizeResponse{}, errors.ErrorNotPng
	}

	readMat := i.gocv.IMRead(retrieveFileRes.FileDestination, gocv.IMReadAnyColor)

	interpolationFlag := gocv.InterpolationFlags(req.InterpolationFlag)

	i.gocv.Resize(readMat, &readMat, image.Point{
		X: req.Width,
		Y: req.Height,
	}, 0, 0, interpolationFlag)

	resizedFileDestination := retrieveFileRes.FileDestination[:len(retrieveFileRes.FileDestination)-4] + "_resized.png"
	isSuccess := i.gocv.IMWrite(resizedFileDestination, readMat)
	if !isSuccess {
		return imageadapter.ResizeResponse{}, errors.ErrorImWrite
	}
	return imageadapter.ResizeResponse{
		FileDestination: resizedFileDestination,
	}, nil
}
