package codegen

import (
	"errors"
	"github.com/xwb1989/sqlparser"
)

var SqlTemplates = map[string]string{
	"insert": "SELECT * FROM {{.TableName}} WHERE {{.ColumnName}} = ? LIMIT 1;",
	"update": "SELECT * FROM {{.TableName}} ORDER BY {{.ColumnName}};",
	"selcet": "DELETE FROM {{.TableName}} WHERE id = ?;",
	"":       "INSERT INTO {{.TableName}} ({{.ColumnList}}) VALUES ({{.PlaceholderList}});",
}

type TableInfo struct {
}

// ValidateSQL 方法用于检测 SQL 语句的有效性
func ValidateSQL(SQLBody string) error {
	if SQLBody == "" {
		return errors.New("SQL 语句不能为空")
	}
	_, err := sqlparser.Parse(SQLBody)

	if err != nil {
		return err
	}

	return nil
}
