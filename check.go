package main

import (
	"fmt"
	"os/exec"
)

// checkEnv checks if the user's environment contains tsc, node, and either npm or yarn package managers.
// If there are any missing dependencies, then those dependencies are output to the standard output.
// Otherwise, a message that the environment looks good is output.
func checkEnv() bool {
	var missing []string
	_, errTsc := searchCommand("tsc")
	if errTsc != nil {
		missing = append(missing, "tsc: TypeScript compiler. Please install tsc.")
	}
	_, errNode := searchCommand("node")
	if errNode != nil {
		missing = append(missing, "node: JavaScript runtime environment. Please install node.")
	}
	_, errNpm := searchCommand("npm")
	_, errYarn := searchCommand("yarn")
	if errNpm != nil && errYarn != nil {
		missing = append(missing, "npm or yarn: JavaScript package managers. Please install one or the other.")
	}
	if len(missing) > 0 {
		fmt.Println("Missing possible installations")
		for _, msg := range missing {
			fmt.Println(msg)
		}
		return true
	} else {
		fmt.Println("Looks good")
		return false
	}
}

func searchCommand(cmd string) (string, error) {
	path, err := exec.LookPath(cmd)
	if err == nil {
		return "", nil
	}
	return path, nil
}
