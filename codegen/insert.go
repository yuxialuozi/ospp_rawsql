package codegen

import (
	"github.com/cloudwego/cwgo/pkg/curd/code"
	"github.com/cloudwego/cwgo/pkg/curd/parse"
)

func insertCodegen(insert *parse.InsertParse) []code.Statement {
	if insert.OperateMode == parse.OperateOne {
		return []code.Statement{
			code.DeclColonStmt{
				Left: code.ListCommaStmt{
					code.RawStmt("result"),
					code.RawStmt("err"),
				},
				Right: code.CallStmt{
					Caller:   code.RawStmt("r.db"),
					CallName: "Create",
					Args: code.ListCommaStmt{
						code.RawStmt(insert.MethodParamNames[1]),
					},
				},
			},
			code.RawStmt("if err != nil {\n\treturn nil, err\n}"),
			code.ReturnStmt{
				ListCommaStmt: code.ListCommaStmt{
					code.RawStmt("result.ID"),
					code.RawStmt("nil"),
				},
			},
		}
	} else {
		return []code.Statement{
			code.DeclColonStmt{
				Left: code.ListCommaStmt{
					code.RawStmt("results"),
					code.RawStmt("err"),
				},
				Right: code.CallStmt{
					Caller:   code.RawStmt("r.db"),
					CallName: "Create",
					Args: code.ListCommaStmt{
						code.RawStmt(insert.MethodParamNames[1]), // Assuming index 1 is slice of users
					},
				},
			},
			code.RawStmt("if err != nil {\n\treturn nil, err\n}"),
			code.ReturnStmt{
				ListCommaStmt: code.ListCommaStmt{
					code.RawStmt("results"), // Assuming Create returns a slice of results
					code.RawStmt("nil"),
				},
			},
		}
	}
}
