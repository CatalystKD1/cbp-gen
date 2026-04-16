package main

import (
	"cbp-gen/gen/c"
	cpp "cbp-gen/gen/cpp"
	gol "cbp-gen/gen/go"
	"cbp-gen/gen/html"
	"flag"
	"fmt"
	"log"
	"strings"
)

func incompatible(badFlag string) {
	err := "You are not allow to use the " + badFlag + " flag here"
	log.Fatal(err)
}

type StringSlice []string

func (s *StringSlice) String() string {
	return strings.Join(*s, ", ")
}

func (s *StringSlice) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func main() {
	// Define programing language flags
	cFlag := flag.Bool("c", false, "Generate a C file")
	cppFlag := flag.Bool("cpp", false, "Generate a C++ file")
	goFlag := flag.Bool("go", false, "Generate a Go file")
	htmlFlag := flag.Bool("html", false, "Generate a HTML file")

	// name flags
	nameFlag := flag.String("name", "main", "Name of the output file")
	titleFlag := flag.String("title", "Hello World", "Title of the HTML page")

	// type flag
	typeFlag := flag.String("type", "int", "Type of the function in the file")

	// include / import flags
	var includes StringSlice
	flag.Var(&includes, "i", "Include/import files")

	// language specific things
	packageFlag := flag.String("pack", "main", "Name of the Go Package")

	var nameSpaces StringSlice
	flag.Var(&nameSpaces, "ns", "Add namespacescd for cpp file")

	flag.Parse()

	if *cFlag {
		c.GenC(*nameFlag, includes, *typeFlag)
	} else if *goFlag {
		gol.GenGo(*nameFlag, includes, *packageFlag)
	} else if *cppFlag {
		cpp.GenCpp(*nameFlag, includes, *typeFlag, nameSpaces)
	} else if *htmlFlag {
		html.GenHtml(*nameFlag, *titleFlag)
	} else {
		fmt.Print("Wrong flag")
	}
}
