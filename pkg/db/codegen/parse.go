package codegen

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type SQLStatement struct {
	Name    string              // 语句名称
	Return  string              // 返回的值，即注释中第二个冒号后面的内容
	APIPath string              // 相关API路径
	SQL     string              // SQL语句内容
	AST     sqlparser.Statement // 解析后的AST节点
}

// processSQLStatements takes a file path, reads the SQL statements, and processes them.
func processSQLStatements(filePath string) (SQLStatement, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sqlStatements := string(data)

	// Split SQL statements by ';'
	statements := strings.Split(sqlStatements, ";")

	// Regex to extract name and API path
	nameAPIRegex := regexp.MustCompile(`--\s*name:\s*(.+?):(\w+)\n--api\.\w+:\s*(.+)`)

	var sqlStatement = SQLStatement{}

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		// Extracting name and API path
		matches := nameAPIRegex.FindStringSubmatch(stmt)
		if len(matches) < 4 {
			log.Fatalf("Error parsing metadata for statement: %s", stmt)
			continue
		}

		name := matches[1]
		execType := matches[2]
		apiPath := matches[3]

		// Extract the SQL part by removing metadata
		sql := nameAPIRegex.ReplaceAllString(stmt, "")

		// Parse the SQL statement
		parsedStmt, err := sqlparser.Parse(sql)
		if err != nil {
			log.Fatalf("Error parsing SQL statement: %s, Error: %v", sql, err)
		}

		// Store in struct
		sqlStatement = SQLStatement{
			Name:    name + ":" + execType,
			Return:  execType,
			APIPath: apiPath,
			SQL:     sql,
			AST:     parsedStmt,
		}

		fmt.Printf("Name: %s\nAPI Path: %s\nSQL:\n%s\n", sqlStatement.Name, sqlStatement.APIPath, sqlStatement.SQL)
		fmt.Printf("Parsed AST:\n%s\n", sqlparser.String(sqlStatement.AST))
		fmt.Println("---------------------------------------------------")
	}

	return sqlStatement, err
}

// sqlTypeToGoType 将 SQL 类型转换为 Go 类型
func sqlTypeToGoType(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "int", "bigint", "smallint", "tinyint":
		return "int64"
	case "varchar", "text", "char":
		return "string"
	case "boolean":
		return "bool"
	default:
		return "interface{}" // 默认情况，使用空接口类型
	}
}
