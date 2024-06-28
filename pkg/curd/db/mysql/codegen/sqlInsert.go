package codegen

import (
	"bytes"
	"fmt"
	"ospp_rawsql/pkg/curd/parse"
	"strings"
)

func sqlInsertCodegen(insert *parse.InsertParse) (string, error) {
	// Generate the INSERT SQL statement
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES ", insert.TableName,
		strings.Join(insert.Columns, ",")))

	// Append values
	for _, row := range insert.Values {
		buf.WriteString("(")
		for i, value := range row {
			if i > 0 {
				buf.WriteString(", ")
			}
			switch v := value.(type) {
			case int, int64, float64:
				buf.WriteString(fmt.Sprintf("%v", v)) // Assuming numeric types
			case string:
				buf.WriteString(fmt.Sprintf("'%s'", strings.Replace(v, "'", "''", -1))) // Escaping single quotes for strings
			default:
				return "", fmt.Errorf("unsupported value type: %T", v)
			}
		}
		buf.WriteString("), ")
	}
	buf.Truncate(buf.Len() - 2) // Remove the last ", "

	// Validate generated SQL statement
	if err := ValidateSQL(buf.String()); err != nil {
		return "", err
	}

	return buf.String(), nil
}
