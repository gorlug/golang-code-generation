package main

import (
	"golang-code-generation/generator"
	"golang-code-generation/todo"
)

func main() {
	parsedStruct := generator.GenerateParsedStruct(todo.TodoEntity{}, "Todo", "todo")
	buffer, err := generator.TemplateToString("struct", parsedStruct)
	if err != nil {
		panic(err)
	}
	formatted, err := generator.FormatResult(buffer)
	if err != nil {
		panic(err)
	}
	err = generator.WriteBytesToFile(formatted, "todo/todo_gen.go")
	if err != nil {
		panic(err)
	}
}
