package schemas

import (
	"io/ioutil"
	"strings"

	"github.com/andrewwormald/gogensql/templates"
)

func Read(path string) ([]templates.Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	schemas := strings.Split(string(b), ";")

	var strcts []templates.Config
	for _, segment := range schemas {
		var strct templates.Config
		nameEl := strings.ReplaceAll(segment, "create table ", "")
		nameEl = strings.Split(nameEl, "(")[0]
		nameEl = strings.TrimSpace(nameEl)
		if nameEl == "" {
			continue
		}

		strct.Table.Name = nameEl
		strctName := snakeToCamelCase(nameEl)
		strct.Struct.Name.Plural = strctName
		strct.Struct.Name.Singular = strings.TrimSuffix(strctName, "s")

		fieldSections := strings.Split(segment, "(")[1]
		fieldSections = strings.ReplaceAll(fieldSections, ")", "")
		fieldSections = strings.TrimSpace(fieldSections)
		fields := strings.Split(fieldSections, ",")

		strct.Struct.Fields = make(map[string]string, len(fields))
		for _, fieldLiteral := range fields {
			fieldLiteral = strings.TrimSpace(fieldLiteral)
			seg := strings.Split(fieldLiteral, " ")
			fieldName := strings.TrimSpace(seg[0])
			if fieldName == "" {
				continue
			}

			strct.Table.Fields = append(strct.Table.Fields, fieldName)
			fieldName = snakeToCamelCase(fieldName)

			typeName := strings.TrimSpace(seg[1])
			if typeName == "" {
				continue
			}

			strct.Struct.Fields[fieldName] = returnGoType(typeName)
		}

		strcts = append(strcts, strct)
	}

	return strcts, nil
}

func returnGoType(sqlType string) string {
	// simplest but messy way of handling for now
 	if strings.Contains(sqlType, "varchar") {
 		return "string"
	}

	switch sqlType {
	case "bigint":
		return "int64"
	case "int":
		return "int32"
	case "tinyint":
		return "bool"
	case "text":
		return "string"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	default:
		return "interface{}"
	}
}

func snakeToCamelCase(in string) string {
	in = strings.TrimSpace(in)
	seg := strings.Split(in, "_")

	newList := make([]string, len(seg))
	for index, char := range seg {
		word := strings.Split(char, "")
		newWord := strings.ToUpper(word[0])

		for i, char := range word {
			if i == 0 {
				continue
			}

			newWord += char
		}

		newList[index] = newWord
	}

	var out string
	for _, char := range newList {
		out += char
	}

	return out
}
