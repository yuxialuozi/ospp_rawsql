package extract

import (
	"github.com/cloudwego/thriftgo/parser"
	"ospp_rawsql/pkg/curd/code"
	"reflect"
)

//这些函数和方法的具体实现可以进一步根据业务逻辑和需求来完善和实现，用于在 Go 语言中提取Thrift IDL 文件和相关任务。

type ThriftUsedInfo struct {
}

func (info *ThriftUsedInfo) ParseThriftIdl() (rawStructs []*IdlExtractStruct, err error) {

	return
}

func extractIdlStruct(st *parser.StructLike, file *parser.Thrift, rawStruct *IdlExtractStruct) error {

	return nil
}

func isThriftBaseType(t string) bool {
	return t == "byte" || t == "i8" || t == "i16" || t == "i32" || t == "i64" ||
		t == "bool" || t == "string" || t == "double" || t == "binary"
}

func isThriftContainerType(t string) bool {
	return t == "map" || t == "set" || t == "list"
}

func convertThriftType(node *parser.Type, file *parser.Thrift) code.Type {

	return nil
}

var thriftBaseTypeMap = map[string]string{
	"byte":   "int8",
	"i8":     "int8",
	"i16":    "int16",
	"i32":    "int32",
	"i64":    "int64",
	"bool":   "bool",
	"string": "string",
	"double": "float64",
}

func handleTagOmitempty(s string) reflect.StructTag {

	return ""
}

func AddMongoModelImports(data string, impt []string) (string, error) {

	return "", nil
}
