package imageservice

import (
	"mini-image-converter/internal/pkg/core/adapter/commonadapter"
	"mini-image-converter/internal/pkg/core/adapter/gocvadapter"
	"mini-image-converter/internal/pkg/core/adapter/osadapter"
)

type ImageService struct {
	gocv   gocvadapter.Adapter
	os     osadapter.Adapter
	common commonadapter.Adapter
}

type NewImageServiceReq struct {
	GoCV   gocvadapter.Adapter
	OS     osadapter.Adapter
	Common commonadapter.Adapter
}

func NewImageService(req NewImageServiceReq) ImageService {
	return ImageService{
		gocv:   req.GoCV,
		os:     req.OS,
		common: req.Common,
	}
}
