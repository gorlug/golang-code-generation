package main

import "golang-code-generation/generator"

type RangeExample struct {
	Name    string
	Hobbies []string
}

func main() {
	buffer, err := generator.TemplateToString("rangeExample", RangeExample{
		Name: "Medium",
		Hobbies: []string{
			"Software Development",
			"Music",
			"Movies",
		},
	})
	if err != nil {
		panic(err)
	}
	println(buffer.String())
	bytesArray := buffer.Bytes()
	err = generator.WriteBytesToFile(&bytesArray, "range-example.txt")
	if err != nil {
		panic(err)
	}
}
