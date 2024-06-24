package codegen

import "github.com/cloudwego/cwgo/pkg/curd/template"

func HandleBaseCodegen() []*template.MethodRender {
	var methods []*template.MethodRender
	methods = append(methods, findOneMethod())
	methods = append(methods, findListMethod())
	methods = append(methods, findPageListMethod())
	methods = append(methods, findSortPageListMethod())
	methods = append(methods, insertOneMethod())
	methods = append(methods, updateOneMethod())
	methods = append(methods, updateManyMethod())
	methods = append(methods, deleteOneMethod())
	methods = append(methods, bulkInsertMethod())
	methods = append(methods, bulkUpdateMethod())
	methods = append(methods, aggregateMethod())
	methods = append(methods, countMethod())

	return methods
}

func countMethod() *template.MethodRender {

	return nil
}

func aggregateMethod() *template.MethodRender {

	return nil
}

func bulkUpdateMethod() *template.MethodRender {
	return nil
}

func bulkInsertMethod() *template.MethodRender {
	return nil
}

func deleteOneMethod() *template.MethodRender {
	return nil
}

func updateManyMethod() *template.MethodRender {
	return nil
}

func updateOneMethod() *template.MethodRender {
	return nil
}

func insertOneMethod() *template.MethodRender {
	return nil
}

func findSortPageListMethod() *template.MethodRender {
	return nil
}

func findPageListMethod() *template.MethodRender {
	return nil
}

func findListMethod() *template.MethodRender {
	return nil
}

func findOneMethod() *template.MethodRender {
	return nil
}
