package tools

import (
	"go/format"
	"os"
)

func WriteFileFormat(fileContent string, path string) {

	formatted, err := format.Source([]byte(fileContent))
	if err != nil {
		panic(err)
	}
	WriteFile(string(formatted), path)

}

func WriteFile(fileContent string, path string) {

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	l2, err := f.WriteString(string(fileContent))
	if err != nil {
		panic(err)
		panic(l2)
		f.Close()
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
