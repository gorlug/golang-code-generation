package generator

import (
	"fmt"
	"reflect"
	"strings"
)

type ParsedStructEnumValue struct {
	Name  string // Name of the enum value
	Value string // The actual value
}

type ParsedStructEnum struct {
	Name   string // Name of the enum to generate
	Values []ParsedStructEnumValue
}

type ParsedStructField struct {
	Name string // Name of the field
	Type string // Type of the field
}

type ParsedStruct struct {
	Name    string              // Name of the struct to generate
	Package string              // Package of the struct to generate
	Fields  []ParsedStructField // Fields of the struct to generate
	Enums   []ParsedStructEnum  // Enums of the struct to generate
}

func GenerateParsedStruct(structType any, structName string, packageName string) ParsedStruct {
	// using reflect we can get meta information about the struct
	t := reflect.TypeOf(structType)
	parsed := ParsedStruct{
		Name:    structName,
		Package: packageName,
		Fields:  []ParsedStructField{},
		Enums:   []ParsedStructEnum{},
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// this is the type of the field, e.g. string, bool, int
		typeName := field.Type.String()
		if isEnum(field) {
			// sets the type name to TodoState which will be generated
			typeName = fmt.Sprint(structName, field.Name)
			// creates a new enum with the name of the field
			parsed.Enums = append(parsed.Enums, ParsedStructEnum{
				Name:   typeName,
				Values: parseEnumValues(getEnumValues(field), structName, field.Name),
			})
		}
		parsed.Fields = append(parsed.Fields, ParsedStructField{
			Name: field.Name, // this is the name of the struct field
			Type: typeName,
		})
	}
	return parsed
}

func parseEnumValues(values []string, structName string, fieldName string) []ParsedStructEnumValue {
	parsedValues := make([]ParsedStructEnumValue, len(values))
	for i, value := range values {
		parsedValues[i] = ParsedStructEnumValue{
			Name:  fmt.Sprint(structName, fieldName, firstLetterToUpper(value)),
			Value: value,
		}
	}
	return parsedValues
}

/**
 * Returns the value of the enum tag split by comma
 */
func getEnumValues(field reflect.StructField) []string {
	enumValues := field.Tag.Get("enum")
	valuesSplit := strings.Split(enumValues, ",")
	return valuesSplit
}

/**
 * Returns true if the field contains the tag enum
 */
func isEnum(field reflect.StructField) bool {
	enumValues := field.Tag.Get("enum")
	return enumValues != ""
}

/**
 * Turns "created" into "Created"
 */
func firstLetterToUpper(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}
