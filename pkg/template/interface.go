package template

import (
	"bytes"
	"ospp_rawsql/pkg/curd/code"
)

var interfaceTemplate = `{{.Comment}}
type {{.Name}} interface {
{{.Methods.GetCode}}
}` + "\n"

type InterfaceRender struct {
	Name    string
	Comment string
	Methods code.InterfaceMethods
}

func (ir *InterfaceRender) RenderObj(buffer *bytes.Buffer) error {
	if err := templateRender(buffer, "interfaceTemplate", interfaceTemplate, ir); err != nil {
		return err
	}
	return nil
}
