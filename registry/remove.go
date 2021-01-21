package registry

import (
	"fmt"
	"os/exec"
)

func Remove(name string) {
	err := exec.Command("yarn", "add", name).Run()
	if err != nil {
		fmt.Println(err)
	}
	err = exec.Command("yarn", "add", "@types/"+name, "--dev").Run()
	if err != nil {
		fmt.Println(err)
	}
}