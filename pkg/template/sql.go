package template

import (
	"bytes"
)

var sqlTemplate = `-- {{.Comment}}
{{.SQLBody}};`

type SQLRender struct {
	Comment string // 注释
	SQLBody string // SQL 语句体
}

func (sr *SQLRender) RenderSQL(buffer *bytes.Buffer) error {
	if _, err := buffer.WriteString(sr.SQLBody + ";\n"); err != nil {
		return err
	}

	return nil
}
