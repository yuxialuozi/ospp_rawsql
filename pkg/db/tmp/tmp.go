package tmp

import (
	"bytes"
	"log"
	"text/template"
)

// SQL 插入模板
var Insert_Template = `
-- name: insert:exec
--api.get: /api/{{.TableName}}/create
INSERT INTO {{.TableName}} ({{range $index, $element := .Fields}}{{if $index}}, {{end}}{{$element}}{{end}})
VALUES ({{range $index, $element := .Fields}}{{if $index}}, {{end}}?{{end}});
`

// InsertRender 用于渲染 SQL 插入模板
type InsertRender struct {
	TableName string   // 表名
	Fields    []string // 字段列表
}

// SQL 根据ID删除模板
var DeleteByIDTemplate = `
-- name: DeleteByID:exec
--api.delete: /api/{{.TableName}}/:id
DELETE FROM {{.TableName}} WHERE Id = ?;
`

// DeleteByIDRender 用于渲染 SQL 根据ID删除模板
type DeleteByIDRender struct {
	TableName string // 表名
}

// SQL 根据ID更新模板
var UpdateByIDTemplate = `
-- name: Update{{.TableName}}ById:exec
--api.put: /api/{{.TableName}}/:id
UPDATE {{.TableName}} SET {{range $index, $field := .Fields}} {{$field}} = ?{{if not (last $index $.Fields)}},{{end}}{{end}}
WHERE Id = ?;
`

// UpdateByIDRender 用于渲染 SQL 根据ID更新模板
type UpdateByIDRender struct {
	TableName string   // 表名
	Fields    []string // 更新字段列表
}

// SQL 获取模板
var GetByTemplate = `
-- name: Get{{.TableName}}By{{.FieldName}}:many
--api.get: /api/{{.TableName}}/{{.FieldName}}/:value
SELECT * FROM {{.TableName}}
WHERE {{.FieldName}} = ?;
`

// GetByRender 用于渲染 SQL 获取模板
type GetByRender struct {
	TableName string // 表名
	FieldName string // 字段名
}

// SQL 删除模板
var DeleteByTemplate = `
-- name: Delete{{.TableName}}By{{.FieldName}}:exec
--api.delete: /api/{{.TableName}}/{{.FieldName}}/:value
DELETE FROM {{.TableName}}
WHERE {{.FieldName}} = ?;
`

// DeleteByRender 用于渲染 SQL 删除模板
type DeleteByRender struct {
	TableName string // 表名
	FieldName string // 字段名
}

// renderTemplate 函数用于渲染模板
func RenderTemplate(templateStr string, data interface{}) string {
	// 注册自定义函数
	funcMap := template.FuncMap{
		"last": func(index int, slice interface{}) bool {
			return index == len(slice.([]string))-1
		},
	}

	// 创建模板并注册函数
	t := template.New("sqlTemplate").Funcs(funcMap)
	tmpl, err := t.Parse(templateStr)
	if err != nil {
		log.Fatalf("解析模板时出错: %v", err)
	}

	// 渲染模板
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		log.Fatalf("渲染模板时出错: %v", err)
	}

	return tpl.String()
}
