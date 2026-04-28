package c

import (
	"embed"
	"os"
	"strings"
	"cbp-gen/models"
)

//go:embed c.tmpl
var templateFiles embed.FS

func GenC(name string, includes []string, typeF string, args []models.FuncArgs) {
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

	argsBlock := ""
	last := args[len(args) - 1]
	for _, arg := range args {
		if (arg.ArgName == last.ArgName) && (arg.ArgType == last.ArgType) {
			argsBlock += arg.ArgType + " " + arg.ArgName
		} else {
			argsBlock += arg.ArgType + " " + arg.ArgName + ", "
		}
	}
	
	result = strings.ReplaceAll(result, "{includes}", includeBlock)
	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{type}", typeF)
	result = strings.ReplaceAll(result, "{args}", argsBlock)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}