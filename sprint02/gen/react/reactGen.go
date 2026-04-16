package react

import (
	"os"
	"strings"
)

// Make a basic react component
// add feature to change the 

func GenGo(name string, hook bool) {
	var output string
	output += name + ".jsx"

	data, err := os.ReadFile("gen/react/component.tmpl")
	if err != nil {
		panic(err)
	}

	

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}