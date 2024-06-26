package codegen

import (
	"ospp_rawsql/pkg/curd/extract"
	"ospp_rawsql/pkg/curd/parse"
	"ospp_rawsql/pkg/template"
)

func HandleCodegen(ifOperations []*parse.InterfaceOperation) (methodRenders [][]*template.MethodRender) {
	for _, ifOperation := range ifOperations {
		methods := make([]*template.MethodRender, 0)
		for _, operation := range ifOperation.Operations {

			switch operation.GetOperationName() {
			case parse.Insert:

			case parse.Find:

			case parse.Update:

			case parse.Delete:

			case parse.Count:

			case parse.Bulk:

			case parse.Transaction:

			default:
			}
		}
		methodRenders = append(methodRenders, methods)
	}
	return
}

var BaseMysqlImports = map[string]string{
	"context": "",
}

func AddMysqlImports(data string) (string, error) {

	return "", nil
}

func GetFuncRender(extractStruct *extract.IdlExtractStruct) *template.FuncRender {

	return nil
}

func GetStructRender(extractStruct *extract.IdlExtractStruct) *template.StructRender {
	return nil
}
