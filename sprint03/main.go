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
	pubFlag := flag.Bool("pub", true, "Used to determine if a function/class is public")

	// include / import flags
	var includes StringSlice
	flag.Var(&includes, "i", "Include/import files")

	// language specific things
	packageFlag := flag.String("pack", "main", "Name of the Go Package")

	var nameSpaces StringSlice
	flag.Var(&nameSpaces, "ns", "Add namespaces for cpp file")

	// function arguments!
	var argList StringSlice
	flag.Var(&argList, "arg", "Add arguments to your function, should follow this format: 'type:name'")

	flag.Parse()

	// --- Validation ---

	// Count how many language flags are set
	langFlags := map[string]bool{
		"c":       *cFlag,
		"cpp":     *cppFlag,
		"go":      *goFlag,
		"html":    *htmlFlag,
		"react-c": *componentFlag,
		"java":    *javaFlag,
		"react":   *reactFlag,
	}

	langCount := 0
	for _, v := range langFlags {
		if v {
			langCount++
		}
	}

	// No language flag — print help and exit
	if langCount == 0 {
		fmt.Println("cbp-gen: A boilerplate code generator")
		fmt.Println("\nUsage: cbp-gen -<language> [options]")
		fmt.Println("\nAvailable flags:")
		flag.PrintDefaults()
		return
	}

	// More than one language flag
	if langCount > 1 {
		log.Fatal("Error: only one language flag can be used at a time")
	}

	// Flags that don't support args
	noArgFlags := map[string]bool{
		"html":  *htmlFlag,
		"react": *reactFlag,
	}
	for name, set := range noArgFlags {
		if set && len(argList) > 0 {
			log.Fatalf("Error: cannot use -arg flag with -%s", name)
		}
	}

	// Flags that don't support includes
	noIncludeFlags := map[string]bool{
		"html":    *htmlFlag,
		"react":   *reactFlag,
		"react-c": *componentFlag,
	}
	for name, set := range noIncludeFlags {
		if set && len(includes) > 0 {
			log.Fatalf("Error: cannot use -i flag with -%s", name)
		}
	}

	// -ts is only for react and react-c
	if *tsFlag && !*reactFlag && !*componentFlag {
		log.Fatal("Error: -ts flag can only be used with -react or -react-c")
	}

	// -ns is only for cpp
	if len(nameSpaces) > 0 && !*cppFlag {
		log.Fatal("Error: -ns flag can only be used with -cpp")
	}

	// -pack is only for go
	if *packageFlag != "main" && !*goFlag {
		log.Fatal("Error: -pack flag can only be used with -go")
	}

	// -title is only for html
	if *titleFlag != "Hello World" && !*htmlFlag {
		log.Fatal("Error: -title flag can only be used with -html")
	}

	// -class is only for java
	if *classFlag != "main" && !*javaFlag {
		log.Fatal("Error: -class flag can only be used with -java")
	}

	// --- Generation ---
	allArgs := getArgs(argList)

	if *cFlag {
		c.GenC(*nameFlag, includes, *typeFlag, allArgs)
	} else if *goFlag {
		gol.GenGo(*nameFlag, includes, *typeFlag, *packageFlag, allArgs)
	} else if *cppFlag {
		cpp.GenCpp(*nameFlag, includes, *typeFlag, nameSpaces, allArgs)
	} else if *htmlFlag {
		html.GenHtml(*nameFlag, *titleFlag)
	} else if *reactFlag {
		react.ReactInit(*nameFlag, *tsFlag)
	} else if *componentFlag {
		react.GenComponent(*nameFlag, *tsFlag, allArgs)
	} else if *javaFlag {
		java.GenJava(*nameFlag, includes, *typeFlag, *pubFlag, *classFlag, allArgs)
	}
}
