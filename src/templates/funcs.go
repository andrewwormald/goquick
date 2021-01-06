package templates

import (
	"os"
	"text/template"
)

var listWhere = `
func list{{.Struct.Name.Plural}}Where(ctx context.Context, dbc *sql.DB, where string) ([]{{.Struct.Name.Singular}}, error) {
	return list{{.Struct.Name.Singular}}AfterFrom(ctx, dbc, where)
}
`

var listAfterFrom = `
func list{{.Struct.Name.Plural}}AfterFrom(ctx context.Context, dbc *sql.DB, afterFromStatement string) ([]{{.Struct.Name.Singular}}, error) {
	rows, err := dbc.QueryContext(ctx, "select {{- range $key, $value := .Table.Fields }}{{if $key}},{{end}} {{$value}}{{end}} from {{.Table.Name}} " + afterFromStatement + ";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ls []{{.Struct.Name.Singular}}
	for rows.Next() {
		var l {{.Struct.Name.Singular}}
		err := rows.Scan(
		{{- range $key, $value := .Struct.Fields }}
			&l.{{ $key }},
		{{- end }}
		)
		if err != nil {
			return nil, err
		}

		ls = append(ls, l)
	}

	return ls, nil
}
`

var lookupWhere = `
func lookup{{.Struct.Name.Singular}}Where(ctx context.Context, dbc *sql.DB, where string) ({{.Struct.Name.Singular}}, error) {
	return lookup{{.Struct.Name.Singular}}AfterFrom(ctx, dbc, where)
}
`

var lookupAfterFrom = `
func lookup{{.Struct.Name.Singular}}AfterFrom(ctx context.Context, dbc *sql.DB, afterFromStatement string) ({{.Struct.Name.Singular}}, error) {
	rows, err := dbc.QueryContext(ctx, "select {{- range $key, $value := .Table.Fields }}{{if $key}},{{end}} {{$value}}{{end}} from {{.Table.Name}} " + afterFromStatement + ";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var l {{.Struct.Name.Singular}}
	for rows.Next() {
		err := rows.Scan(
		{{- range $key, $value := .Struct.Fields }}
			&l.{{ $key }},
		{{- end }}
		)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}
`

var insert = `
func Insert{{.Struct.Name.Singular}}(ctx context.Context, dbc *sql.DB, strct {{.Struct.Name.Singular}}) (int64, error) {
	res, err := dbc.ExecContext(ctx, "insert into {{.Table.Name}} "+
		"set {{- range $key, $value := .Table.Fields }} {{if $key}},{{end}}{{$value}}=?
		{{- end }}",
		{{- range $key, $value := .Struct.Fields }}
			strct.{{ $key }},
		{{- end }}
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
`

var delete = `
func Delete{{.Struct.Name.Singular}}(ctx context.Context, dbc *sql.DB, where string) error {
	_, err := dbc.ExecContext(ctx, "delete from {{.Table.Name}} "+ where)
	return err
}
`

var update = `
func Update{{.Struct.Name.Singular}}(ctx context.Context, dbc *sql.DB, where, set string, args... interface{}) error {
	_, err = dbc.ExecContext(ctx, "update {{.Table.Name}} set " + set + where, args)
	return err
}
`


func createFunctionSet(file *os.File, in Config) error {
	templates := []string{listAfterFrom, listWhere, lookupAfterFrom, lookupWhere, insert, delete, update}
	for _, temp := range templates {
		err := template.Must(template.New("").Parse(temp)).Execute(file, in)
		if err != nil {
			return err
		}
	}

	return nil
}
