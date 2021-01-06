package templates

import (
	"os"
	"text/template"
)

func writeStruct(file *os.File, in Config) error {
	return template.Must(template.New("").Parse(StructTemplate)).Execute(file, in)
}

type Config struct {
	PackageName string
	Struct      struct {
		Name struct {
			Singular string
			Plural   string
		}
		Fields map[string]string
	}
	Table struct {
		Name   string
		Fields []string
	}
}

var StructTemplate = `
type {{.Struct.Name.Singular}} struct {
{{- range $key, $value := .Struct.Fields }}
	{{ $key }} {{ $value }}
{{- end }}
}
`
