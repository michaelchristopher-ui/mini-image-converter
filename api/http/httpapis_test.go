package http_test

import (
	"mime/multipart"
	apihttp "mini-image-converter/api/http"
	"mini-image-converter/internal/pkg/core/adapter/imageadapter"
	"mini-image-converter/internal/pkg/platform/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestConvert(t *testing.T) {
	mockImageService, mockEcho, apiIntegrator := InitTest(t)
	tests := map[string]struct {
		mock func()
	}{
		"is return the correct error when formfile return error": {
			mock: func() {
				mockEcho.EXPECT().FormFile(mock.Anything).Once().Return(nil, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorFileParam,
				}).Once().Return(nil)
			},
		},
		"is return the correct error when convert returns error": {
			mock: func() {
				mockEcho.EXPECT().FormFile(mock.Anything).Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Convert(mock.Anything).Once().Return(imageadapter.ConvertResponse{}, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusInternalServerError, apihttp.ErrorJson{
					Error: apihttp.ErrorInternalServer,
				}).Once().Return(nil)
			},
		},
		"is executing file function when all is good": {
			mock: func() {
				mockEcho.EXPECT().FormFile(mock.Anything).Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Convert(mock.Anything).Once().Return(imageadapter.ConvertResponse{}, nil)
				mockEcho.EXPECT().File("").Once().Return(nil)
				mockImageService.EXPECT().RemoveFile(mock.Anything).Once().Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			apiIntegrator.Convert(mockEcho)
		})
	}
}

func TestResize(t *testing.T) {
	mockImageService, mockEcho, apiIntegrator := InitTest(t)
	tests := map[string]struct {
		mock func()
	}{
		"is returning the correct error when width is not valid": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("a")
				mockEcho.EXPECT().FormValue("height").Once().Return("b")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("c")
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorWidthParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when height is not valid": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("b")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("c")
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorHeightParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when interpolation flag is not valid due to it not being an integer string": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("c")
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorInterpolationFlagParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when interpolation flag is not valid due to it not being more than 0": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("-1")
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorInterpolationFlagParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when interpolation flag is not valid due to it not being less than 7": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("8")
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorInterpolationFlagParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when form file is not valid": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("1")
				mockEcho.EXPECT().FormFile("file").Once().Return(nil, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorFileParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when resize fails": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("1")
				mockEcho.EXPECT().FormFile("file").Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Resize(mock.Anything).Once().Return(imageadapter.ResizeResponse{}, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusInternalServerError, apihttp.ErrorJson{
					Error: apihttp.ErrorInternalServer,
				}).Once().Return(nil)
			},
		},
		"is executing file function when all is good": {
			mock: func() {
				mockEcho.EXPECT().FormValue("width").Once().Return("1")
				mockEcho.EXPECT().FormValue("height").Once().Return("1")
				mockEcho.EXPECT().FormValue("interpolation_flag").Once().Return("1")
				mockEcho.EXPECT().FormFile("file").Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Resize(mock.Anything).Once().Return(imageadapter.ResizeResponse{}, nil)
				mockEcho.EXPECT().File("").Once().Return(nil)
				mockImageService.EXPECT().RemoveFile("").Once().Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			apiIntegrator.Resize(mockEcho)
		})
	}
}

func TestCompress(t *testing.T) {
	mockImageService, mockEcho, apiIntegrator := InitTest(t)
	tests := map[string]struct {
		mock func()
	}{
		"is returning the correct error when formfile returns an error": {
			mock: func() {
				mockEcho.EXPECT().FormFile("file").Once().Return(&multipart.FileHeader{}, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusBadRequest, apihttp.ErrorJson{
					Error: apihttp.ErrorFileParam,
				}).Once().Return(nil)
			},
		},
		"is returning the correct error when resize returns an error": {
			mock: func() {
				mockEcho.EXPECT().FormFile("file").Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Compress(mock.Anything).Once().Return(imageadapter.CompressResponse{}, errors.ErrFoo)
				mockEcho.EXPECT().JSON(http.StatusInternalServerError, apihttp.ErrorJson{
					Error: apihttp.ErrorInternalServer,
				}).Once().Return(nil)
			},
		},
		"is executing file function when all is good": {
			mock: func() {
				mockEcho.EXPECT().FormFile("file").Once().Return(&multipart.FileHeader{}, nil)
				mockImageService.EXPECT().Compress(mock.Anything).Once().Return(imageadapter.CompressResponse{}, nil)
				mockEcho.EXPECT().File("").Once().Return(nil)
				mockImageService.EXPECT().RemoveFile("").Once().Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			apiIntegrator.Compress(mockEcho)
		})
	}
}
