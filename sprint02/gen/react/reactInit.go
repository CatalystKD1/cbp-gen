package react

import (
	"os/exec"
)

func ReactInit(name string, ts bool) {
	var cmd *exec.Cmd
	if ts {
		cmd = exec.Command("npm", "create", "vite@latest", name, "--", "--template", "react-ts")
	} else {
		cmd = exec.Command("npm", "create", "vite@latest", name, "--", "--template", "react")
	}
	
	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
}