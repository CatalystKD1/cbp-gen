package main

import (
	//"fmt"
	"os"
	"log"
	"os/exec"
)

const filePath = "test/subdir/test.txt"

func main() {
	// make a directory
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("filePath", []byte("Hello, Gophers!"), 0600)
	if err != nil {
		log.Fatal(err)
	}
}