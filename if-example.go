package main

import "golang-code-generation/generator"

type IfExample struct {
	Name     string
	ShowMore bool
}

func main() {
	generateTemplate(IfExample{
		Name:     "Medium",
		ShowMore: false,
	}, "range-example-false.txt")

	generateTemplate(IfExample{
		Name:     "Medium",
		ShowMore: true,
	}, "range-example-true.txt")
}

func generateTemplate(falseExample IfExample, path string) {
	buffer, err := generator.TemplateToString("ifExample", falseExample)
	if err != nil {
		panic(err)
	}
	println(buffer.String())
	bytesArray := buffer.Bytes()
	err = generator.WriteBytesToFile(&bytesArray, path)
	if err != nil {
		panic(err)
	}
}
