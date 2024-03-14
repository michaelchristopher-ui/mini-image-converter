package ioadapter

import "io"

type Adapter interface {
	Copy(dst io.Writer, src io.Reader) (written int64, err error)
}
