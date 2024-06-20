package codegen

import (
	"fmt"
	"strings"
)

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
