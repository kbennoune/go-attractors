package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func fileWriter(filename string) (func(calculation), func()) {
	path := filepath.Dir(filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	fo, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	finalizer := func() {
		fmt.Println("Closing ", filename)
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}

	writer := func(calc calculation) {
		writeFile(calc, *fo)
	}

	return writer, finalizer
}

func writeFile(c calculation, f os.File) {
	output := fmt.Sprintf("%d %f %f\n", c.last.t, c.last.x, c.last.y)
	_, err := f.WriteString(output)

	if err != nil {
		panic(err)
	}
}
