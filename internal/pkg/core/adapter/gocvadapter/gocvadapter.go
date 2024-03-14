package gocvadapter

import (
	"image"

	"gocv.io/x/gocv"
)

type Adapter interface {
	IMWriteWithParams(name string, img gocv.Mat, params []int) bool
	IMRead(name string, flags gocv.IMReadFlag) gocv.Mat
	IMWrite(name string, img gocv.Mat) bool
	Resize(src gocv.Mat, dst *gocv.Mat, sz image.Point, fx float64, fy float64, interp gocv.InterpolationFlags)
}
