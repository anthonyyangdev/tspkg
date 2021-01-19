package main

import (
	"fmt"
	"os"
	"os/exec"

	"./registry"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Make sure to include the command")
		return
	}
	command := args[0]

	switch command {
	case "check":
		checkEnv()
	case "init":
		initProject()
	case "add":
		if len(args) < 1 {
			fmt.Println("Make sure to include what package to add")
			return
		}
		target := args[1]
		registry.Get(target)
	case "remove":
	default:
		fmt.Println("Command", command, "is not recognized")
	}
}

func initProject() {
	err := exec.Command("tsc", "--init").Run()
	if err != nil {
		panic(err)
	}
	buildPackage()
}
