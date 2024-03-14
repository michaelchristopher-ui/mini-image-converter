package http

type ErrorJson struct {
	Error string `json:"error"`
}

const (
	ErrorWidthParam             = "invalid width parameter"
	ErrorHeightParam            = "invalid height parameter"
	ErrorInterpolationFlagParam = "invalid interpolation flag parameter"
	ErrorFileParam              = "missing file parameter"
	ErrorInternalServer         = "internal server error"
)
