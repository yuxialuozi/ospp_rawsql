package main

import (
	"fmt"
	"os"
)

func main() {
	operations := []tmplate.Operation{
		{Name: "InsertOne", Type: "execresult"},
		{Name: "InsertMany", Type: "execresult"},
		{Name: "FindUsernames", Type: "many"},
		{Name: "FindByUsernameAge", Type: "one"},
		{Name: "UpdateContact", Type: "exec"},
		{Name: "DeleteById", Type: "exec"},
		{Name: "CountByAge", Type: "one"},
	}

	tableInfo := tmplate.TableInfo{
		TableName:       "users",
		ColumnName:      "id",
		ColumnList:      "username, email, age",
		PlaceholderList: "?, ?, ?",
	}

	// Create sql folder if not exists
	err := os.MkdirAll("sql", os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating sql folder: %s\n", err)
		return
	}

	fileName := "sql/user.sql"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fileName, err)
		return
	}
	defer file.Close()

	for _, op := range operations {
		sqlTemplate, found := tmplate.SqlTemplates[op.Type]
		if !found {
			fmt.Printf("Template not found for operation: %s\n", op.Name)
			continue
		}

		tmpl, err := template.New(op.Name).Parse(sqlTemplate)
		if err != nil {
			fmt.Printf("Error parsing template for operation %s: %s\n", op.Name, err)
			continue
		}

		_, err = file.WriteString(fmt.Sprintf("-- name: %s :%s\n", op.Name, op.Type))
		if err != nil {
			fmt.Printf("Error writing to file %s: %s\n", fileName, err)
			return
		}

		err = tmpl.Execute(file, tableInfo)
		if err != nil {
			fmt.Printf("Error executing template for operation %s: %s\n", op.Name, err)
			continue
		}

		fmt.Printf("Generated SQL for operation %s\n", op.Name)
	}
	fmt.Printf("Generated SQL file: %s\n", fileName)
}
