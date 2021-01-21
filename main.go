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
		for _, pkg := range args[1:] {
			registry.Get(pkg)
		}
	case "remove":
		for _, pkg := range args[1:] {
			registry.Remove(pkg)
		}
	default:
		fmt.Println("Command", command, "is not recognized")
	}
}

func initProject() {
	err := buildPackage()
	if err != nil {panic(err)}
	err = exec.Command("tsc", "--init").Run()
	if err != nil {panic(err)}
}
