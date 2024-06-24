package codegen

import "ospp_rawsql/pkg/curd/code"

func getCtx() code.Param {
	return code.Param{
		Name: "ctx",
		Type: code.IdentType("context.Context"),
	}
}

func GetMysqlInsertOneParams() (res []code.Param) {
	ctx := getCtx()
	insertOneData := code.Param{
		Name: "insertOneData",
		Type: code.InterfaceType{
			Name: "interface{}",
		},
	}
	res = append(res, ctx, insertOneData)
	return
}

func GetMDeleteOneParams() (res []code.Param) {
	ctx := getCtx()
	deleteOneData := code.Param{
		Name: "deleteOneData",
		Type: code.InterfaceType{
			Name: "interface{}",
		},
	}
	res = append(res, ctx, deleteOneData)
	return
}

func GetMFindOneParams() (res []code.Param) {
	ctx := getCtx()
	selector := code.Param{
		Name: "selector",
		Type: code.IdentType("bson.M"),
	}
	result := code.Param{
		Name: "result",
		Type: code.InterfaceType{
			Name: "interface{}",
		},
	}
	res = append(res, ctx, selector, result)
	return
}

func GetMBulkInsertParams() (res []code.Param) {
	ctx := getCtx()
	batchData := code.Param{
		Name: "batchData",
		Type: code.IdentType("[]interface{}"),
	}
	res = append(res, ctx, batchData)
	return
}

func GetMBulkUpdateParams() (res []code.Param) {
	ctx := getCtx()
	filter := code.Param{
		Name: "filter",
		Type: code.IdentType("[]interface{}"),
	}
	updater := code.Param{
		Name: "updater",
		Type: code.IdentType("[]interface{}"),
	}
	res = append(res, ctx, filter, updater)
	return
}

func GetMAggregateParams() (res []code.Param) {
	ctx := getCtx()
	pipeline := code.Param{
		Name: "pipeline",
		Type: code.IdentType("[]bson.M"),
	}
	result := code.Param{
		Name: "result",
		Type: code.IdentType("interface{}"),
	}
	res = append(res, ctx, pipeline, result)
	return
}

func GetMCountParams() (res []code.Param) {
	ctx := getCtx()
	selector := code.Param{
		Name: "selector",
		Type: code.IdentType("bson.M"),
	}
	res = append(res, ctx, selector)
	return
}
