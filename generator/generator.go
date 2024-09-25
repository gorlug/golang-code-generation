package generator

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"
)

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

func FormatResult(buffer *bytes.Buffer) (*[]byte, error) {
	formattedBytes, err := format.Source(buffer.Bytes())
	if err != nil {
		return &[]byte{}, err
	}
	return &formattedBytes, nil
}
