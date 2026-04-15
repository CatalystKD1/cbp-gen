package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()  
	fmt.Println("current Dir: ", dir)

	fullPath := filepath.Join(dir, "main.go")
	// get specific components of the dir
	fmt.Println("File Name (Base):", filepath.Base(fullPath))
	fmt.Println("Directory (Dir):", filepath.Dir(fullPath))
	fmt.Println("Extension:", filepath.Ext(fullPath))
}