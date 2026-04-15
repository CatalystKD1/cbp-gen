package c

import (
	"os"
	"strings"
)

// test variables
var name string
var includes []string

func GenC() {
	name = "main"
	includes = append(includes, "stdio.h")
	includes = append(includes, "math.h")


	data, err := os.ReadFile("gen/c/c.tmpl")
	if err != nil {
		panic(err)
	}

	result := string(data)
	// Build the full includes string first
	includeBlock := ""
	for _, include := range includes {
		includeBlock += "#include <" + include + ">\n"
	}
	result = strings.ReplaceAll(result, "{includes}", includeBlock)

	result = strings.ReplaceAll(result, "{name}", name)

	err = os.WriteFile("output.c", []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}