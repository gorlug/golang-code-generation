package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type Example struct {
	Name string
}

func main() {
	buffer, err := TemplateToString("simpleExample", Example{Name: "World"})
	if err != nil {
		panic(err)
	}
	println(buffer.String())
	bytesArray := buffer.Bytes()
	err = WriteBytesToFile(&bytesArray, "simple-example.txt")
	if err != nil {
		panic(err)
	}
}

func TemplateToString(templateName string, data any) (*bytes.Buffer, error) {
	tmpl := template.Must(template.ParseGlob("generator/template/*.tmpl"))
	buffer := bytes.NewBuffer([]byte{})
	err := tmpl.ExecuteTemplate(buffer, templateName, data)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func WriteBytesToFile(bytes *[]byte, path string) error {
	outputFile, err := os.Create(path)
	if err != nil {
		log.Printf("failed to create file: %v", err)
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.Write(*bytes)
	if err != nil {
		log.Printf("failed to write file: %v", err)
		return err
	}
	return nil
}
