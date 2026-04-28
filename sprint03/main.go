package main

import (
	c "cbp-gen/gen/c"
	cpp "cbp-gen/gen/cpp"
	gol "cbp-gen/gen/go"
	"cbp-gen/gen/html"
	"cbp-gen/gen/java"
	"cbp-gen/gen/react"
	"cbp-gen/models"
	"flag"
	"fmt"
	"log"
	"strings"
)



func getArgs(args []string) []models.FuncArgs {
	// this function will take a slice of string, then return 
	// a list of function arguments. First it will need to slice the string aroung the colon :
	// format of a proper models.FuncArgs "type:name" split the string based on the colon. if there is no colon then 
	// exit the program with na error using log.Fatel(err)

	result := []models.FuncArgs{}

	for _, arg := range args {
		parts := strings.Split(arg, ":") // split into 2 elements in a slice ["type", "name"]
		if len(parts) != 2 { // check if there are 2 elements (intended behaviour)
			log.Fatalf("Invalid argument format %q: expected 'type:name'", arg)
		}

		result = append(result, models.FuncArgs {
			ArgType: parts[0],
			ArgName: parts[1],
		})
	}
	return result
}

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
	componentFlag := flag.Bool("react-c", false, "Generate a React Component (jsx) file")
	javaFlag := flag.Bool("java", false, "Generate a Java file")
	/*kotlinFlag
	rustFlag
	phpFlag
	csharpFlag*/

	// name flags
	nameFlag := flag.String("name", "main", "Name of the output file")
	titleFlag := flag.String("title", "Hello World", "Title of the HTML page")
	classFlag := flag.String("class", "main", "The name of a class (useful for Java and C#)")

	// type flag
	typeFlag := flag.String("type", "int", "Type of the function in the file")
	tsFlag := flag.Bool("ts", false, "Starts a TypeScript React document")
	reactFlag := flag.Bool("react", false, "Start a React project using npm")
	pubFlag := flag.Bool("pub", false, "Used to determine if a function/class is public")

	// include / import flags
	var includes StringSlice
	flag.Var(&includes, "i", "Include/import files")

	// language specific things
	packageFlag := flag.String("pack", "main", "Name of the Go Package")

	var nameSpaces StringSlice
	flag.Var(&nameSpaces, "ns", "Add namespacescd for cpp file")

	// function arguments!
	var argList StringSlice
	flag.Var(&argList, "arg", "Add arguments to your function, should follow this format: 'type:name'")
	
	// new section
	flag.Parse()

	// get all args and deals with them
	allArgs := getArgs(argList)

	if *cFlag {
		c.GenC(*nameFlag, includes, *typeFlag, allArgs)
	} else if *goFlag {
		gol.GenGo(*nameFlag, includes, *typeFlag, *packageFlag)
	} else if *cppFlag {
		cpp.GenCpp(*nameFlag, includes, *typeFlag, nameSpaces)
	} else if *htmlFlag {
		html.GenHtml(*nameFlag, *titleFlag)
	} else if *reactFlag {
		react.ReactInit(*nameFlag, *tsFlag)
	} else if *componentFlag {
		react.GenComponent(*nameFlag, *tsFlag)
	} else if *javaFlag {
		java.GenJava(*nameFlag, includes, *typeFlag, *pubFlag, *classFlag)
	} else {
		fmt.Print("Please write a flag for a programming language.")
	}
}
