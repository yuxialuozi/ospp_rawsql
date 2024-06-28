package codegen

import "ospp_rawsql/pkg/curd/code"

func insertOneBaseCodegen() []code.Statement {
	stmt := `if insertOneData == nil {
		return nil, fmt.Errorf("insert param is empty")
	}

	return b.collection.InsertOne(ctx, insertOneData)`

	return []code.Statement{
		code.RawStmt(stmt),
	}
}
