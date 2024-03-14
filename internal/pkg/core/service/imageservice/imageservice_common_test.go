package imageservice_test

import (
	"mini-image-converter/internal/pkg/core/service/imageservice"
	"mini-image-converter/internal/pkg/mocks/mockcommonadapter"
	"mini-image-converter/internal/pkg/mocks/mockgocvadapter"
	"mini-image-converter/internal/pkg/mocks/mockosadapter"
	"testing"
)

type InitTestRet struct {
	Is         imageservice.ImageService
	MockGoCV   *mockgocvadapter.Adapter
	MockOS     *mockosadapter.Adapter
	MockCommon *mockcommonadapter.Adapter
}

func InitTest(t *testing.T) (*mockgocvadapter.Adapter,
	*mockosadapter.Adapter, *mockcommonadapter.Adapter,
	imageservice.ImageService) {
	mockGoCV := mockgocvadapter.NewAdapter(t)
	mockOS := mockosadapter.NewAdapter(t)
	mockCommon := mockcommonadapter.NewAdapter(t)
	is := imageservice.NewImageService(imageservice.NewImageServiceReq{
		GoCV:   mockGoCV,
		OS:     mockOS,
		Common: mockCommon,
	})
	return mockGoCV, mockOS, mockCommon, is
}
