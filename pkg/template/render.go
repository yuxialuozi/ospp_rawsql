package template

import "bytes"

type Render interface {
	RenderObj(buffer *bytes.Buffer) error
}
