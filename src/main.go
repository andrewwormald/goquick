// GoQuick enables faster development for simple processes such as basic web server storage

package main

import (
	"flag"
	"github.com/andrewwormald/gogensql/templates"
	"os"
	"strings"

	"github.com/andrewwormald/gogensql/schemas"
)

var schemaFilePath = flag.String("schema_path", "", "The relative path to the schema.sql file used for generation")
var packageName = flag.String("package_name", "", "The package that the generated file will be part of")
var tableNames = flag.String("tables", "", "csv list of tables that must only be generated and not the whole file")

func main() {
	flag.Parse()

	// get the current location that its being executed from
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// change directory to its being executed from
	err = os.Chdir(wd)
	if err != nil {
		panic(err)
	}

	// now can use provided path as relative to that of the go generate command
	templs, err := schemas.Read(*schemaFilePath)
	if err != nil {
		panic(err)
	}

	if *tableNames != "" {
		templs = filter(*tableNames, templs)
	}

	templates.Generate(*packageName, templs)
}

func filter(tableNames string, tmpls []templates.Config) []templates.Config {
	var filtered []templates.Config
	names := strings.Split(tableNames, ",")
	for _, t := range tmpls {
		for _, nme := range names {
			if t.Table.Name == nme {
				filtered = append(filtered, t)
			}
		}
	}

	return filtered
}
