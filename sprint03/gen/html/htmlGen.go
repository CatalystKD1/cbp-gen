package html

import (
	"os"
	"strings"
	"embed"
)

var templateFiles embed.FS

// test variables
/*var name string
var includes []string*/


func GenHtml(name string, title string) {
	var output string
	if (name == "main") {
		name = "index"
	}
	output += name + ".html"

	data, err :=  templateFiles.ReadFile("hhtml.tmpl")
	if err != nil {
		panic(err)
	}
	title = "<title>\n\t" + "\t" + "\t" + title + "\n\t" + "\t" + "</title>"

	result := string(data)
	result = strings.ReplaceAll(result, "{title}", title)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}