package template

import (
	"bytes"

	"github.com/cloudwego/cwgo/pkg/curd/code"
)

var funcTemplate = `{{.Comment}}
func {{.Name}}{{.Params.GetCode}} {{.Returns.GetCode}} {
{{.FuncBody.GetCode}}
}` + "\n"

type FuncRender struct {
	Name     string
	Comment  string
	Params   code.Params
	Returns  code.Returns
	FuncBody code.Body
}

func (fr *FuncRender) RenderObj(buffer *bytes.Buffer) error {
	if err := templateRender(buffer, "funcTemplate", funcTemplate, fr); err != nil {
		return err
	}
	return nil
}
