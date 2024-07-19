package thrift

// 数据库字段属性和约束定义
const (
	Type                    = "type="
	ForeignKey              = "foreign_key="
	PrimaryKey              = "primary_key"
	AutoIncrement           = "auto_increment"
	DefaultValue            = "default_value="
	NotNull                 = "NOT NULL"
	Unique                  = "UNIQUE"
	Check                   = "CHECK"
	Index                   = "INDEX"
	DefaultCurrentTimestamp = "DEFAULT CURRENT_TIMESTAMP"
	Cascade                 = "CASCADE"
	Comment                 = "COMMENT"
)

const (
	AnnotationQuery    = "api.query"
	AnnotationForm     = "api.form"
	AnnotationPath     = "api.path"
	AnnotationHeader   = "api.header"
	AnnotationCookie   = "api.cookie"
	AnnotationBody     = "api.body"
	AnnotationRawBody  = "api.raw_body"
	AnnotationJsConv   = "api.js_conv"
	AnnotationNone     = "api.none"
	AnnotationFileName = "api.file_name"
)
