package tmplate

var SqlTemplates = map[string]string{
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

type Operation struct {
	Name    string
	Type    string
	SQLType string
}
