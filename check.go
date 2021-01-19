package main

import (
	"fmt"
	"os/exec"
)

func checkEnv() {
	missing := searchCommand("tsc")
	missing = append(missing, searchCommand("node")...)
	if len(missing) > 0 {
		fmt.Println("Missing possible installations", missing)
	} else {
		fmt.Println("Looks good")
	}
}

func searchCommand(cmd string) []string {
	_, err := exec.LookPath(cmd)
	if err == nil {
		return []string{}
	}
	return []string{cmd}
}
