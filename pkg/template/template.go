package template

import "bytes"

type Template struct {
	Renders []Render
}

func (t *Template) AddRender(render Render) {
	t.Renders = append(t.Renders, render)
}

func (t *Template) Build() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	for _, render := range t.Renders {
		if err := render.RenderObj(buffer); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}
