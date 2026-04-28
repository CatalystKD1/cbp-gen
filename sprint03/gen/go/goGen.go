package gol

import (
	"cbp-gen/models"
	"embed"
	"os"
	"strings"
)

var templateFiles embed.FS

// test variables
/*var name string
var includes []string*/


func GenGo(name string, includes []string, typeF string, packages string, args []models.FuncArgs) {
	var output string
	output += name + ".go"

	data, err :=  templateFiles.ReadFile("go.tmpl")
	if err != nil {
		panic(err)
	}

	/*
		Make explicite void functions
	*/

	result := string(data)
	// Build the full includes string first
	includeBlock := ""
	for _, include := range includes {
		includeBlock += "\"" + include + "\"\n"
	}

	if (name == "main" && typeF == "int") {
		typeF = ""
	}

	argsBlock := ""
	last := args[len(args) - 1]
	for _, arg := range args {
		if (arg.ArgName == last.ArgName) && (arg.ArgType == last.ArgType) {
			argsBlock += arg.ArgName + " " + arg.ArgType
		} else {
			argsBlock += arg.ArgName + " " + arg.ArgType + ", "
		}
	}

	result = strings.ReplaceAll(result, "{includes}", includeBlock)
	result = strings.ReplaceAll(result, "{type}", typeF)
	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{package}", packages)
	result = strings.ReplaceAll(result, "{args}", argsBlock)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}