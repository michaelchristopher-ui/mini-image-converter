package http_test

import (
	apihttp "mini-image-converter/api/http"
	"mini-image-converter/internal/pkg/mocks/mockecho"
	"mini-image-converter/internal/pkg/mocks/mockimageadapter"
	"testing"
)

func InitTest(t *testing.T) (*mockimageadapter.Adapter, *mockecho.Context, *apihttp.APIIntegrator) {
	mockImageService := mockimageadapter.NewAdapter(t)
	mockEcho := mockecho.NewContext(t)
	apiIntegrator := apihttp.NewAPIIntegrator(mockImageService)
	return mockImageService, mockEcho, apiIntegrator
}
