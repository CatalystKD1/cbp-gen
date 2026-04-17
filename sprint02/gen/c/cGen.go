package c

import (
	"embed"
	"os"
	"strings"
)

//go:embed c.tmpl
var templateFiles embed.FS

func GenC(name string, includes []string, typeF string) {
	var output string
	output += name + ".c"

	data, err := templateFiles.ReadFile("c.tmpl")
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
	result = strings.ReplaceAll(result, "{type}", typeF)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}