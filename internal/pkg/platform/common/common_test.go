package common_test

import (
	"bytes"
	"mini-image-converter/internal/pkg/core/adapter/commonadapter"
	"mini-image-converter/internal/pkg/mocks/mockcommonadapter"
	"mini-image-converter/internal/pkg/mocks/mockioadapter"
	"mini-image-converter/internal/pkg/mocks/mockosadapter"
	"mini-image-converter/internal/pkg/mocks/mockother"
	"mini-image-converter/internal/pkg/platform/common"
	"testing"

	"mini-image-converter/internal/pkg/platform/errors"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type MockMultipartFile struct {
	*bytes.Reader
}

func (m *MockMultipartFile) Close() error {
	return nil
}

func NewMockMultipartFile(content []byte) *MockMultipartFile {
	return &MockMultipartFile{
		Reader: bytes.NewReader(content),
	}
}

func TestRetrieveFile(t *testing.T) {
	mockOS := mockosadapter.NewAdapter(t)
	mockFile := mockcommonadapter.NewFileOpener(t)
	mockIO := mockioadapter.NewAdapter(t)
	commonx := common.NewCommon(mockIO)

	tests := map[string]struct {
		mock    func()
		withRes commonadapter.RetrieveFileRes
		withErr error
	}{
		"is return correct error when opening file": {
			mock: func() {
				mockFile.EXPECT().Open().Once().Return(nil, errors.ErrFoo)
			},
			withErr: errors.ErrFoo,
		},
		"is return correct error when fail to create": {
			mock: func() {
				mockFile.EXPECT().Open().Once().Return(NewMockMultipartFile([]byte("mock file content")), nil)
				mockOS.EXPECT().Create(mock.Anything).Once().Return(nil, errors.ErrFoo)
			},
			withErr: errors.ErrFoo,
		},
		"is return correct error after io copy": {
			mock: func() {
				mockFile.EXPECT().Open().Once().Return(NewMockMultipartFile([]byte("mock file content")), nil)
				mockWriteCloser := mockother.NewWriteCloser(t)
				mockOS.EXPECT().Create(mock.Anything).Once().Return(mockWriteCloser, nil)
				mockWriteCloser.EXPECT().Close().Once().Return(nil)
				mockIO.EXPECT().Copy(mock.Anything, mock.Anything).Once().Return(int64(0), errors.ErrFoo)
			},
			withErr: errors.ErrFoo,
		},
		"is return correct response": {
			mock: func() {
				mockFile.EXPECT().Open().Once().Return(NewMockMultipartFile([]byte("mock file content")), nil)
				mockWriteCloser := mockother.NewWriteCloser(t)
				mockOS.EXPECT().Create(mock.Anything).Once().Return(mockWriteCloser, nil)
				mockWriteCloser.EXPECT().Close().Once().Return(nil)
				mockIO.EXPECT().Copy(mock.Anything, mock.Anything).Once().Return(int64(0), nil)
			},
			withRes: commonadapter.RetrieveFileRes{
				FileDestination: "dummy.png",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			res, err := commonx.RetrieveFile(mockFile, "dummy.png", mockOS)
			assert.Equal(t, res, tc.withRes)
			assert.Equal(t, err, tc.withErr)
		})
	}
}

func TestIsPNG(t *testing.T) {
	mockIO := mockioadapter.NewAdapter(t)
	commonx := common.NewCommon(mockIO)
	os := mockosadapter.NewAdapter(t)

	tests := map[string]struct {
		mock    func()
		name    string
		withRes bool
		withErr error
	}{
		"is error file reader": {
			mock: func() {
				os.EXPECT().ReadFile(mock.Anything).Once().Return([]byte{}, errors.ErrFoo)
			},
			withRes: false,
			withErr: errors.ErrFoo,
		},
		"is valid png": {
			mock: func() {
				os.EXPECT().ReadFile(mock.Anything).Once().Return([]byte("\x89PNG\r\n\x1a\n"), nil)
			},
			withRes: true,
		},
		"is invalid png": {
			mock: func() {
				os.EXPECT().ReadFile(mock.Anything).Once().Return([]byte("NOTAPNG"), nil)
			},
			withRes: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock()
			result, err := commonx.IsPNG(os, "dummy.png")
			assert.Equal(t, err, tc.withErr)
			assert.Equal(t, result, tc.withRes)
		})
	}
}
