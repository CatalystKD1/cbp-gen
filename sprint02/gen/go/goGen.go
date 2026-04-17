package gol

import (
	"os"
	"strings"
	"embed"
)

var templateFiles embed.FS

// test variables
/*var name string
var includes []string*/


func GenGo(name string, includes []string, packages string) {
	var output string
	output += name + ".go"

	data, err :=  templateFiles.ReadFile("go.tmpl")
	if err != nil {
		panic(err)
	}

	result := string(data)
	// Build the full includes string first
	includeBlock := ""
	for _, include := range includes {
		includeBlock += "\"" + include + "\"\n"
	}
	result = strings.ReplaceAll(result, "{includes}", includeBlock)

	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{package}", packages)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}