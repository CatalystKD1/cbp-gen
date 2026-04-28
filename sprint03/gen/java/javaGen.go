package java

import (
	"cbp-gen/models"
	"embed"
	"os"
	"strings"
)

//go:embed java.tmpl
var templateFiles embed.FS

func GenJava(name string, includes []string, typeF string, publicF bool, class string, args []models.FuncArgs) {
	var output string
	output += name + ".java"

	data, err :=  templateFiles.ReadFile("java.tmpl")
	if err != nil {
		panic(err)
	}

	/*
		Make explicite void functions
	*/
	public := ""

	if publicF {
		public = "public"
	}

	if (name == "main") {
		name = "Main"
		if len(args) == 0 {
			// this slice is empty
			arg := models.FuncArgs {
				ArgType: "String[]",
				ArgName: "args",
			}
			args = append(args, arg)
		}
	}

	result := string(data)
	// Build the full includes string first
	includeBlock := ""
	for _, include := range includes {
		includeBlock += "import" + include + "\n"
	}

	if (name == "main" && typeF == "int") {
		typeF = "void"
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
	result = strings.ReplaceAll(result, "{type}", typeF)
	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{public}", public)
	result = strings.ReplaceAll(result, "{class}", class)
	result = strings.ReplaceAll(result, "{args}", argsBlock)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}