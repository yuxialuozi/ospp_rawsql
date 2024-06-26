package template

import (
	"bytes"

	"github.com/cloudwego/cwgo/pkg/curd/code"
)

var methodTemplate = `{{.Comment}}
func {{.MethodReceiver.GetCode}} {{.Name}}{{.Params.GetCode}} {{.Returns.GetCode}} {
{{.MethodBody.GetCode}}
}` + "\n"

type MethodRender struct {
	Name           string
	Comment        string
	MethodReceiver code.MethodReceiver
	Params         code.Params
	Returns        code.Returns
	MethodBody     code.Body
}

func (mr *MethodRender) RenderObj(buffer *bytes.Buffer) error {
	if err := templateRender(buffer, "methodTemplate", methodTemplate, mr); err != nil {
		return err
	}
	return nil
}
