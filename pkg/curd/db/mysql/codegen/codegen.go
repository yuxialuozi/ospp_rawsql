package codegen

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"github.com/cloudwego/cwgo/pkg/curd/code"
	"github.com/cloudwego/cwgo/pkg/curd/extract"
	"github.com/cloudwego/cwgo/pkg/curd/parse"
	"github.com/cloudwego/cwgo/pkg/curd/template"

	"golang.org/x/tools/go/ast/astutil"
)

func HandleCodegen(ifOperations []*parse.InterfaceOperation) (methodRenders [][]*template.MethodRender) {
	for _, ifOperation := range ifOperations {
		methods := make([]*template.MethodRender, 0)
		for _, operation := range ifOperation.Operations {

			switch operation.GetOperationName() {
			case parse.Insert:
				insert := operation.(*parse.InsertParse)
				method := &template.MethodRender{
					Name: insert.BelongedToMethod.Name,
					MethodReceiver: code.MethodReceiver{
						Name: "r",
						Type: code.StarExprType{
							RealType: code.IdentType(ifOperation.BelongedToStruct.Name + "RepositoryMongo"),
						},
					},
					Params:     insert.BelongedToMethod.Params,
					Returns:    insert.BelongedToMethod.Returns,
					MethodBody: insertCodegen(insert),
				}
				methods = append(methods, method)

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

var BaseMongoImports = map[string]string{
	"context": "",
}

func AddMongoImports(data string) (string, error) {
	fSet := token.NewFileSet()
	file, err := parser.ParseFile(fSet, "", data, parser.ParseComments)
	if err != nil {
		return "", err
	}

	flagBson, flagMongo, flagOption := false, false, false
	ast.Inspect(file, func(n ast.Node) bool {
		if importSpec, ok := n.(*ast.ImportSpec); ok && importSpec.Path.Value == "go.mongodb.org/mongo-driver/bson" {
			flagBson = true
			return false
		}
		if importSpec, ok := n.(*ast.ImportSpec); ok && importSpec.Path.Value == "go.mongodb.org/mongo-driver/mongo" {
			flagMongo = true
			return false
		}
		if importSpec, ok := n.(*ast.ImportSpec); ok && importSpec.Path.Value == "go.mongodb.org/mongo-driver/mongo/options" {
			flagOption = true
			return false
		}
		return true
	})

	if strings.Contains(data, "bson") {
		if !flagBson {
			astutil.AddNamedImport(fSet, file, "", "go.mongodb.org/mongo-driver/bson")
		}
	}
	if strings.Contains(data, "mongo") {
		if !flagMongo {
			astutil.AddNamedImport(fSet, file, "", "go.mongodb.org/mongo-driver/mongo")
		}
	}
	if strings.Contains(data, "options") {
		if !flagOption {
			astutil.AddNamedImport(fSet, file, "", "go.mongodb.org/mongo-driver/mongo/options")
		}
	}

	buf := new(bytes.Buffer)
	if err = printer.Fprint(buf, fSet, file); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetFuncRender(extractStruct *extract.IdlExtractStruct) *template.FuncRender {
	return &template.FuncRender{
		Name: "New" + extractStruct.Name + "Repository",
		Params: code.Params{
			code.Param{
				Name: "collection",
				Type: code.StarExprType{
					RealType: code.SelectorExprType{
						X:   "mongo",
						Sel: "Collection",
					},
				},
			},
		},
		Returns: code.Returns{
			code.IdentType(extractStruct.Name + "Repository"),
		},
		FuncBody: code.Body{
			code.RawStmt("return &" + extractStruct.Name + "RepositoryMongo{\n\tcollection: collection,\n}"),
		},
	}
}

func GetStructRender(extractStruct *extract.IdlExtractStruct) *template.StructRender {
	return &template.StructRender{
		Name: extractStruct.Name + "RepositoryMongo",
		StructFields: code.StructFields{
			code.StructField{
				Name: "collection",
				Type: code.StarExprType{
					RealType: code.SelectorExprType{
						X:   "mongo",
						Sel: "Collection",
					},
				},
			},
		},
	}
}
