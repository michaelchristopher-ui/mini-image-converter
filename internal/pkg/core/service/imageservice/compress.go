package imageservice

import (
	"mini-image-converter/internal/pkg/core/adapter/imageadapter"
	"mini-image-converter/internal/pkg/platform/errors"
	"os"

	"gocv.io/x/gocv"
)

func (i ImageService) Compress(req imageadapter.CompressRequest) (imageadapter.CompressResponse, error) {
	retrieveFileRes, err := i.common.RetrieveFile(req.File, req.FileName, i.os)
	if err != nil {
		return imageadapter.CompressResponse{}, err
	}
	defer i.os.Remove(retrieveFileRes.FileDestination)

	//Check Content Type
	isPng, err := i.common.IsPNG(i.os, retrieveFileRes.FileDestination)
	if err != nil {
		return imageadapter.CompressResponse{}, err
	}
	if !isPng {
		return imageadapter.CompressResponse{}, errors.ErrorNotPng
	}

	readMat := i.gocv.IMRead(retrieveFileRes.FileDestination, gocv.IMReadAnyColor)

	imWriteParams := []int{
		gocv.IMWritePngCompression,
		9,
	}

	compressedFileDestination := retrieveFileRes.FileDestination[:len(retrieveFileRes.FileDestination)-4] + "_compressed.png"
	isSuccess := i.gocv.IMWriteWithParams(compressedFileDestination, readMat, imWriteParams)
	if !isSuccess {
		os.Remove(retrieveFileRes.FileDestination)
		return imageadapter.CompressResponse{}, errors.ErrorImWrite
	}
	return imageadapter.CompressResponse{
		FileDestination: compressedFileDestination,
	}, nil
}
