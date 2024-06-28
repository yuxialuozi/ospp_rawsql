package template

import (
	"bytes"
	"fmt"
	"text/template"
)

var baseTemplate = `// Code generated by cwgo ({{.Version}}). DO NOT EDIT.

package {{.PackageName}}

import (
{{.GetImports}}
)` + "\n"

type BaseRender struct {
	Version     string            // cwgo version
	PackageName string            // package name in target generation go file
	Imports     map[string]string // key:import path value:import name
}

func (bt *BaseRender) RenderObj(buffer *bytes.Buffer) error {
	if err := templateRender(buffer, "baseTemplate", baseTemplate, bt); err != nil {
		return err
	}
	return nil
}

func (bt *BaseRender) GetImports() string {
	result := ""

	for key, value := range bt.Imports {
		if value != "" {
			result += fmt.Sprintf("\tvalue "+`"%s"`+"\n", key)
		} else {
			result += fmt.Sprintf("\t"+`"%s"`+"\n", key)
		}
	}

	return result
}

func templateRender(buffer *bytes.Buffer, templateName, parseText string, data any) error {
	tmpl, err := template.New(templateName).Parse(parseText)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(buffer, data); err != nil {
		return err
	}

	return nil
}