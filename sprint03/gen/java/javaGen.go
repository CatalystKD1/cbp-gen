package java

import (
	"os"
	"strings"
	"embed"
)

var templateFiles embed.FS

func GenJava(name string, includes []string, typeF string, publicF bool, class string) {
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

	result := string(data)
	// Build the full includes string first
	includeBlock := ""
	for _, include := range includes {
		includeBlock += "import" + include + "\n"
	}

	if (name == "main" && typeF == "int") {
		typeF = "void"
	}

	result = strings.ReplaceAll(result, "{includes}", includeBlock)
	result = strings.ReplaceAll(result, "{type}", typeF)
	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{public}", public)
	result = strings.ReplaceAll(result, "{class}", class)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}