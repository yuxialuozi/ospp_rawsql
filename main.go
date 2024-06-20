package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Operation struct {
	Name    string
	Type    string
	SQLType string
}

var sqlTemplates = map[string]string{
	"one":        "SELECT * FROM {{.TableName}} WHERE {{.ColumnName}} = ? LIMIT 1;",
	"many":       "SELECT * FROM {{.TableName}} ORDER BY {{.ColumnName}};",
	"exec":       "DELETE FROM {{.TableName}} WHERE id = ?;",
	"execresult": "INSERT INTO {{.TableName}} ({{.ColumnList}}) VALUES ({{.PlaceholderList}});",
}

type TableInfo struct {
	TableName       string
	ColumnName      string
	ColumnList      string
	PlaceholderList string
}

// queryName := "FindByLbLbUsernameEqualOrUsernameEqualRbAndAgeGreaterThanRb"
func generateSQL(queryName string) string {
	parts := strings.Split(queryName, "By")
	if len(parts) != 2 {
		return ""
	}

	// 解析操作类型
	operationType := parts[0]

	// 解析条件
	condition := parts[1]
	conditionParts := strings.Split(condition, "And")
	if len(conditionParts) == 0 {
		return ""
	}

	// 构建 WHERE 子句
	whereClause := "WHERE "
	for _, part := range conditionParts {
		equalParts := strings.Split(part, "Equal")
		if len(equalParts) != 2 {
			continue
		}
		fieldName := equalParts[0]
		whereClause += fmt.Sprintf("%s = ? AND ", fieldName)
	}

	// 去除末尾的 " AND "
	whereClause = strings.TrimSuffix(whereClause, " AND ")

	// 构建 SQL 查询语句
	var sql string
	switch operationType {
	case "Find":
		sql = fmt.Sprintf("SELECT * FROM users %s LIMIT 1;", whereClause)
	case "Delete":
		sql = fmt.Sprintf("DELETE FROM users %s;", whereClause)
	case "Count":
		sql = fmt.Sprintf("SELECT COUNT(*) FROM users %s;", whereClause)
	default:
		// 如果操作类型未知，则返回空字符串
		sql = ""
	}
	return sql
}

func main() {
	operations := []Operation{
		{Name: "InsertOne", Type: "execresult"},
		{Name: "InsertMany", Type: "execresult"},
		{Name: "FindUsernames", Type: "many"},
		{Name: "FindByUsernameAge", Type: "one"},
		{Name: "UpdateContact", Type: "exec"},
		{Name: "DeleteById", Type: "exec"},
		{Name: "CountByAge", Type: "one"},
	}

	tableInfo := TableInfo{
		TableName:       "users",
		ColumnName:      "id",
		ColumnList:      "username, email, age",
		PlaceholderList: "?, ?, ?",
	}

	// Create sql folder if not exists
	err := os.MkdirAll("sql", os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating sql folder: %s\n", err)
		return
	}

	fileName := "sql/user.sql"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fileName, err)
		return
	}
	defer file.Close()

	for _, op := range operations {
		sqlTemplate, found := sqlTemplates[op.Type]
		if !found {
			fmt.Printf("Template not found for operation: %s\n", op.Name)
			continue
		}

		tmpl, err := template.New(op.Name).Parse(sqlTemplate)
		if err != nil {
			fmt.Printf("Error parsing template for operation %s: %s\n", op.Name, err)
			continue
		}

		_, err = file.WriteString(fmt.Sprintf("-- name: %s :%s\n", op.Name, op.Type))
		if err != nil {
			fmt.Printf("Error writing to file %s: %s\n", fileName, err)
			return
		}

		err = tmpl.Execute(file, tableInfo)
		if err != nil {
			fmt.Printf("Error executing template for operation %s: %s\n", op.Name, err)
			continue
		}

		fmt.Printf("Generated SQL for operation %s\n", op.Name)
	}
	fmt.Printf("Generated SQL file: %s\n", fileName)
}
