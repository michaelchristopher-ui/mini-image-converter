package imageadapter

import "mime/multipart"

type Adapter interface {
	Resize(req ResizeRequest) (ResizeResponse, error)
	Convert(req ConvertRequest) (ConvertResponse, error)
	Compress(req CompressRequest) (CompressResponse, error)
	RemoveFile(fileName string)
}

type ResizeRequest struct {
	File              *multipart.FileHeader
	FileName          string
	Width             int
	Height            int
	InterpolationFlag int
}

type ResizeResponse struct {
	FileDestination string
}

type ConvertRequest struct {
	File     *multipart.FileHeader
	FileName string
}

type ConvertResponse struct {
	JpegName string
}

type CompressRequest struct {
	File     *multipart.FileHeader
	FileName string
}

type CompressResponse struct {
	FileDestination string
}
