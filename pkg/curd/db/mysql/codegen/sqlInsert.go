package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"ospp_rawsql/pkg/curd/parse"
	"text/template"
)

func sqlInsertCodegen(insert *parse.InsertParse) (string, error) {
	// 确定模板类型
	var templateKey string

	switch insert.OperateMode {
	case parse.OperateOne:
		templateKey = "one"
	case parse.OperateMany:
		templateKey = "many"
	default:
		return "", errors.New("unsupported operate mode")
	}

	// 获取对应的模板
	tmplStr, ok := SqlTemplates[templateKey]
	if !ok {
		return "", fmt.Errorf("template not found for key: %s", templateKey)
	}

	// 准备模板
	tmpl, err := template.New("sqlTemplate").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	// 准备数据
	data := TableInfo{
		TableName:       insert.BelongedToMethod.Name,
		ColumnName:      insert.MethodParamNames[0],
		ColumnList:      "column1, column2, column3", // 替换为实际列名列表
		PlaceholderList: "?, ?, ?",                   // 替换为实际占位符列表
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	// 验证生成的 SQL 语句
	if err := ValidateSQL(buf.String()); err != nil {
		return "", err
	}

	return buf.String(), nil
}
