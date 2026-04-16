package cpp

import (
	"os"
	"strings"
	"fmt"
)

// test variables
/*var name string
var includes []string*/


func GenCpp(name string, includes []string, typeF string, nameSpaces []string) {

	var output string
	output += name + ".cpp"

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

	nameSpaceBlock := ""
	for _, nameSpace := range nameSpaces {
		nameSpaceBlock += "using namespace " + nameSpace + ";\n"
	}

	result = strings.ReplaceAll(result, "{includes}", includeBlock)
	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{type}", typeF)
	result = strings.ReplaceAll(result, "{namespace}", nameSpaceBlock)

	
	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}