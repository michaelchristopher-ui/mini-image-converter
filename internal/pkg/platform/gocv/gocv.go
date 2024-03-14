package gocv

import (
	"image"

	"gocv.io/x/gocv"
)

type GoCV struct {
}

func NewGoCV() GoCV {
	return GoCV{}
}

func (g GoCV) IMWriteWithParams(name string, img gocv.Mat, params []int) bool {
	return gocv.IMWriteWithParams(name, img, params)
}
func (g GoCV) IMRead(name string, flags gocv.IMReadFlag) gocv.Mat {
	return gocv.IMRead(name, flags)
}
func (g GoCV) IMWrite(name string, img gocv.Mat) bool {
	return gocv.IMWrite(name, img)
}
func (g GoCV) Resize(src gocv.Mat, dst *gocv.Mat, sz image.Point, fx float64, fy float64, interp gocv.InterpolationFlags) {
	gocv.Resize(src, dst, sz, fx, fy, interp)
}
