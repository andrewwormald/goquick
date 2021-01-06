package templates

import (
	"go/format"
	"io/ioutil"
	"os"
)

func Generate(filename string, in []Config) error {
	f, err := os.Create(filename + "_gen.go")
	if err != nil {
		return err
	}

	err = writePackageHeader(f, PackageHeader{Name: filename})
	if err != nil {
		return err
	}

	for _, config := range in {
		err := writeStruct(f, config)
		if err != nil {
			return err
		}

		err = createFunctionSet(f, config)
	}

	f.Close()

	b, err := ioutil.ReadFile(filename + "_gen.go")
	if err != nil {
		return err
	}

	formatted, err := format.Source(b)
	if err != nil {
		return err
	}

	f, err = os.Create(filename + "_gen.go")
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		return err
	}

	return nil
}
