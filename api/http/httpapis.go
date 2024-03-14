package http

import (
	"net/http"
	"strconv"

	"mini-image-converter/internal/pkg/core/adapter/imageadapter"

	"github.com/labstack/echo"
)

type APIIntegrator struct {
	ImageService imageadapter.Adapter
}

func NewAPIIntegrator(imageService imageadapter.Adapter) *APIIntegrator {
	return &APIIntegrator{ImageService: imageService}
}

func API(e *echo.Echo, imageService imageadapter.Adapter) {
	api := e.Group("")
	integrator := NewAPIIntegrator(imageService)

	api.POST("/resize", integrator.Resize)
	api.POST("/convert", integrator.Convert)
	api.POST("/compress", integrator.Compress)
}

func (integrator *APIIntegrator) Convert(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorFileParam,
		})
	}

	res, err := integrator.ImageService.Convert(imageadapter.ConvertRequest{
		File:     file,
		FileName: file.Filename,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorJson{
			Error: ErrorInternalServer,
		})
	}

	defer integrator.ImageService.RemoveFile(res.JpegName)
	return c.File(res.JpegName)
}

func (integrator *APIIntegrator) Resize(c echo.Context) error {
	widthString := c.FormValue("width")
	heightString := c.FormValue("height")
	interpolationFlagString := c.FormValue("interpolation_flag")

	width, err := strconv.Atoi(widthString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorWidthParam,
		})
	}

	height, err := strconv.Atoi(heightString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorHeightParam,
		})
	}

	interpolationFlagInt, err := strconv.Atoi(interpolationFlagString)
	if err != nil || (interpolationFlagInt < 0 || (interpolationFlagInt > 4 && interpolationFlagInt != 7)) {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorInterpolationFlagParam,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorFileParam,
		})
	}

	res, err := integrator.ImageService.Resize(imageadapter.ResizeRequest{
		Width:             width,
		Height:            height,
		File:              file,
		FileName:          file.Filename,
		InterpolationFlag: interpolationFlagInt,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorJson{
			Error: ErrorInternalServer,
		})
	}

	defer integrator.ImageService.RemoveFile(res.FileDestination)
	return c.File(res.FileDestination)
}

func (integrator *APIIntegrator) Compress(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJson{
			Error: ErrorFileParam,
		})
	}

	res, err := integrator.ImageService.Compress(imageadapter.CompressRequest{
		File:     file,
		FileName: file.Filename,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorJson{
			Error: ErrorInternalServer,
		})
	}

	defer integrator.ImageService.RemoveFile(res.FileDestination)
	return c.File(res.FileDestination)
}
