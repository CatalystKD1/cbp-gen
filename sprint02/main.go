package main

import (
	"fmt"
	"cbp-gen/gen/c"
)

func main() {
	c.GenC()
	fmt.Print("Check your file")
}