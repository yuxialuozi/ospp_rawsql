package thrift

import (
	"fmt"
	"github.com/cloudwego/thriftgo/parser"
	"ospp_rawsql/pkg/db/tmp"
	"regexp"
	"strings"
)

// GenerateSQLFromAST 根据 AST 生成 SQL 代码
func GenerateSQLFromAST(ast *parser.Thrift) (string, error) {
	if ast == nil {
		return "", fmt.Errorf("AST is nil")
	}

	var sb strings.Builder

	for _, strct := range ast.Structs {
		tableName, err := extractTableName(strct.ReservedComments)
		if err != nil {
			return "", err
		}

		sb.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))
		for i, field := range strct.Fields {
			sqlType, tags := parseTags(field)

			// 处理每个 tag
			sqlType, tags = handleTags(sqlType, tags)

			// 添加字段定义，并将 tag 注释添加到字段定义中
			sb.WriteString(fmt.Sprintf("  %s %s", field.Name, sqlType))
			if len(tags) > 0 {
				sb.WriteString(fmt.Sprintf("  %s ", strings.Join(tags, " ")))
			}
			if len(field.Annotations) > 1 {
				// 将 Values 转换为字符串
				values := strings.Join(field.Annotations[1].Values, ",")
				sb.WriteString(fmt.Sprintf(" --%s:%s", field.Annotations[1].Key, values))
			}

			if i < len(strct.Fields)-1 {
				sb.WriteString(",\n")
			}
		}
		sb.WriteString("\n);\n\n")
	}

	return sb.String(), nil
}

// extractTableName 从注释字符串中提取表名
func extractTableName(comments string) (string, error) {
	// 编写正则表达式来匹配表名注释
	tableNameRe := regexp.MustCompile(`\/\/table_name:(\w+)`)
	match := tableNameRe.FindStringSubmatch(comments)
	if len(match) > 1 {
		return match[1], nil
	}
	return "", fmt.Errorf("struct contains no table_name")
}

// extractTags 从字段的注释中提取 tags
func extractTags(field *parser.Field) []string {
	var tags []string
	for _, annotation := range field.Annotations {
		if annotation.Key == "idl.tag" && len(annotation.Values) > 0 {
			tags = strings.Split(annotation.Values[0], ",")
			break
		}
	}
	return tags
}

// parseTags 解析字段的 idl.tag 注释，返回 SQL 类型和其他标签注释
func parseTags(field *parser.Field) (string, []string) {
	tags := extractTags(field)

	var sqlType string

	// 解析 type 标签
	for _, tag := range tags {
		tag = strings.TrimSpace(tag) // 去除两端空格

		// 检查 type 标签
		if strings.HasPrefix(tag, Type) {
			sqlType = strings.TrimPrefix(tag, Type)
			sqlType = strings.TrimSpace(sqlType)
		}
	}

	sqlType = convertType(field.Type)

	return sqlType, tags
}

// handleTags 处理每个 tag 并返回更新后的 SQL 类型和标签注释
func handleTags(sqlType string, tags []string) (string, []string) {
	var updatedTags []string

	for _, tag := range tags {
		tag = strings.TrimSpace(tag) // 去除两端空格

		switch {
		case strings.HasPrefix(tag, PrimaryKey):
			// 处理 primary_key 标签
			sqlType += " PRIMARY KEY"
		case strings.HasPrefix(tag, AutoIncrement):
			// 处理 auto_increment 标签
			sqlType += " AUTO_INCREMENT"
		case strings.HasPrefix(tag, ForeignKey):
			// 处理 foreign_key 标签
			foreignKey := strings.TrimPrefix(tag, ForeignKey)
			updatedTags = append(updatedTags, fmt.Sprintf("FOREIGN KEY %s", foreignKey))
		case strings.HasPrefix(tag, DefaultValue):
			// 处理 default_value 标签
			defaultValue := strings.TrimPrefix(tag, DefaultValue)
			sqlType += fmt.Sprintf(" DEFAULT %s", defaultValue)
		case strings.HasPrefix(tag, NotNull):
			// 处理 NOT NULL 标签
			sqlType += " NOT NULL"
		case strings.HasPrefix(tag, Unique):
			// 处理 UNIQUE 标签
			sqlType += " UNIQUE"
		case strings.HasPrefix(tag, Check):
			// 处理 CHECK 标签
			check := strings.TrimPrefix(tag, Check)
			updatedTags = append(updatedTags, fmt.Sprintf("CHECK %s", check))
		case strings.HasPrefix(tag, Index):
			// 处理 INDEX 标签
			index := strings.TrimPrefix(tag, Index)
			updatedTags = append(updatedTags, fmt.Sprintf("INDEX %s", index))
		case strings.HasPrefix(tag, DefaultCurrentTimestamp):
			// 处理 DEFAULT CURRENT_TIMESTAMP 标签
			sqlType += " DEFAULT CURRENT_TIMESTAMP"
		case strings.HasPrefix(tag, Cascade):
			// 处理 CASCADE 标签
			sqlType += " CASCADE"
		case strings.HasPrefix(tag, Comment):
			// 处理 COMMENT 标签
			comment := strings.TrimPrefix(tag, Comment)
			updatedTags = append(updatedTags, fmt.Sprintf("COMMENT '%s'", comment))
		case strings.HasPrefix(tag, Type):
			continue
		default:
			// 其他标签
			updatedTags = append(updatedTags, tag)
		}
	}

	// 将 tags 中的逗号替换为空格
	for i, tag := range updatedTags {
		updatedTags[i] = strings.ReplaceAll(tag, ",", " ")
	}

	return sqlType, updatedTags
}

// convertType 将 Thrift 类型转换为 SQL 类型
func convertType(t *parser.Type) string {
	switch t.Name {
	case "i64":
		return "BIGINT"
	case "string":
		return "VARCHAR(255)"
	case "i32":
		return "INT"
	case "bool":
		return "BOOLEAN"
	default:
		return "VARCHAR(255)"
	}
}

func GenerateSQLQuerisFromAST(ast *parser.Thrift) (string, error) {
	if ast == nil {
		return "", fmt.Errorf("AST is nil")
	}

	var builder strings.Builder
	for _, strct := range ast.Structs {
		sqlQueries, err := generateSQLQueriesForStruct(strct)
		if err != nil {
			return "", err
		}
		builder.WriteString(sqlQueries)
	}

	return builder.String(), nil
}

func generateSQLQueriesForStruct(strct *parser.StructLike) (string, error) {
	tableName, err := extractTableName(strct.ReservedComments)
	if err != nil {
		return "", err
	}
	var builder strings.Builder

	// 插入语句
	insertRender := tmp.InsertRender{
		TableName: tableName,
		Fields:    getFieldNames(strct.Fields),
	}
	insertQuery := tmp.RenderTemplate(tmp.Insert_Template, insertRender)
	builder.WriteString(insertQuery)
	builder.WriteString("\n")

	// 创建根据ID删除语句
	deleteByIDRender := tmp.DeleteByIDRender{
		TableName: tableName,
	}
	deleteByIDQuery := tmp.RenderTemplate(tmp.DeleteByIDTemplate, deleteByIDRender)
	builder.WriteString(deleteByIDQuery)
	builder.WriteString("\n")

	// 创建根据ID更新语句
	updateByIDRender := tmp.UpdateByIDRender{
		TableName: tableName,
		Fields:    getFieldNames(strct.Fields),
	}
	updateByIDQuery := tmp.RenderTemplate(tmp.UpdateByIDTemplate, updateByIDRender)
	builder.WriteString(updateByIDQuery)
	builder.WriteString("\n")

	// 遍历每个字段生成查询和删除语句
	for _, field := range strct.Fields {
		// 创建查询语句
		getByRender := tmp.GetByRender{
			TableName: tableName,
			FieldName: field.Name,
		}
		getByQuery := tmp.RenderTemplate(tmp.GetByTemplate, getByRender)
		builder.WriteString(getByQuery)
		builder.WriteString("\n")

		// 创建根据字段名删除语句
		deleteByRender := tmp.DeleteByRender{
			TableName: tableName,
			FieldName: field.Name,
		}
		deleteByQuery := tmp.RenderTemplate(tmp.DeleteByTemplate, deleteByRender)
		builder.WriteString(deleteByQuery)
		builder.WriteString("\n")
	}

	return builder.String(), nil
}

func getFieldNames(fields []*parser.Field) []string {
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = f.Name
	}
	return names
}
