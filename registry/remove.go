package registry

import (
	"fmt"
	"os/exec"
)

func Remove(name string) {
	err := exec.Command("yarn", "remove", name).Run()
	if err != nil {
		fmt.Println(err)
	}
	err = exec.Command("yarn", "remove", "@types/"+name).Run()
	if err != nil {
		fmt.Println(err)
	}
}