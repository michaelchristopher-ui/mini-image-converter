package imageservice_test

import (
	"mini-image-converter/internal/pkg/core/adapter/commonadapter"
	"mini-image-converter/internal/pkg/core/adapter/imageadapter"
	internalError "mini-image-converter/internal/pkg/platform/errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"gocv.io/x/gocv"
	"gotest.tools/assert"
)

func TestResize(t *testing.T) {
	mockGoCV, mockOS, mockCommon, is := InitTest(t)

	tests := map[string]struct {
		mock     func()
		req      imageadapter.ResizeRequest
		wantResp imageadapter.ResizeResponse
		wantErr  error
	}{
		"is returning the correct error when retrieve file returns an error": {
			mock: func() {
				mockCommon.EXPECT().RetrieveFile(mock.Anything, mock.Anything, mock.Anything).Once().Return(commonadapter.RetrieveFileRes{}, errFoo)
			},
			req:      imageadapter.ResizeRequest{},
			wantResp: imageadapter.ResizeResponse{},
			wantErr:  errFoo,
		},
		"is returning the correct error when checking png": {
			mock: func() {
				mockCommon.EXPECT().RetrieveFile(mock.Anything, mock.Anything, mock.Anything).Once().Return(commonadapter.RetrieveFileRes{}, nil)
				mockCommon.EXPECT().IsPNG(mock.Anything, mock.Anything).Once().Return(false, errFoo)
				mockOS.EXPECT().Remove(mock.Anything).Once().Return(nil)
			},
			req:      imageadapter.ResizeRequest{},
			wantResp: imageadapter.ResizeResponse{},
			wantErr:  errFoo,
		},
		"is returning the correct error as if png is not supplied": {
			mock: func() {
				mockCommon.EXPECT().RetrieveFile(mock.Anything, mock.Anything, mock.Anything).Once().Return(commonadapter.RetrieveFileRes{}, nil)
				mockCommon.EXPECT().IsPNG(mock.Anything, mock.Anything).Once().Return(false, nil)
				mockOS.EXPECT().Remove(mock.Anything).Once().Return(nil)
			},
			req:      imageadapter.ResizeRequest{},
			wantResp: imageadapter.ResizeResponse{},
			wantErr:  internalError.ErrorNotPng,
		},
		"is returning the correct error when imwrite fails": {
			mock: func() {
				mockCommon.EXPECT().RetrieveFile(mock.Anything, mock.Anything, mock.Anything).Once().Return(commonadapter.RetrieveFileRes{
					FileDestination: "x.png",
				}, nil)
				mockCommon.EXPECT().IsPNG(mock.Anything, mock.Anything).Once().Return(true, nil)
				mockGoCV.EXPECT().IMRead(mock.Anything, mock.Anything).Once().Return(gocv.Mat{})
				mockGoCV.EXPECT().Resize(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				mockGoCV.EXPECT().IMWrite(mock.Anything, mock.Anything).Once().Return(false)
				mockOS.EXPECT().Remove(mock.Anything).Once().Return(nil)
			},
			req:      imageadapter.ResizeRequest{},
			wantResp: imageadapter.ResizeResponse{},
			wantErr:  internalError.ErrorImWrite,
		},
		"is success resize": {
			mock: func() {
				mockCommon.EXPECT().RetrieveFile(mock.Anything, mock.Anything, mock.Anything).Once().Return(commonadapter.RetrieveFileRes{
					FileDestination: "x.png",
				}, nil)
				mockCommon.EXPECT().IsPNG(mock.Anything, mock.Anything).Once().Return(true, nil)
				mockGoCV.EXPECT().IMRead(mock.Anything, mock.Anything).Once().Return(gocv.Mat{})
				mockGoCV.EXPECT().IMWrite(mock.Anything, mock.Anything).Once().Return(true)
				mockOS.EXPECT().Remove(mock.Anything).Once().Return(nil)
			},
			req: imageadapter.ResizeRequest{},
			wantResp: imageadapter.ResizeResponse{
				FileDestination: "x_resized.png",
			},
			wantErr: nil,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			res, err := is.Resize(tc.req)
			assert.Equal(t, res, tc.wantResp)
			assert.Equal(t, err, tc.wantErr)
		})
	}
}
