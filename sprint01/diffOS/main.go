package main

import (
	"runtime"
	"fmt"
	//"os"
)

// want to detect which OS I am using, then change the file path to match the system
func main() {
	switch i := runtime.GOOS; i {
	case "darwin": // mac os
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	}
}