package template

import (
	"bytes"

	"github.com/cloudwego/cwgo/pkg/curd/code"
)

var structTemplate = `{{.Comment}}
type {{.Name}} struct {
{{.StructFields.GetCode}}
}` + "\n"

type StructRender struct {
	Name         string
	Comment      string
	StructFields code.StructFields
}

func (sr *StructRender) RenderObj(buffer *bytes.Buffer) error {
	if err := templateRender(buffer, "structTemplate", structTemplate, sr); err != nil {
		return err
	}
	return nil
}
