package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
)

func main() {
	// initialized a mode, but checks if it already exists
	 _, err := os.Stat("go.mod") 
	if os.IsNotExist(err) {
		cmd := exec.Command("go", "mod", "init", "test")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Initialized a module")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File exists")
	}
	
}