package registry

import (
	"fmt"
	"net/http"
	"os/exec"

	"../api"
)

func Get(name string) {
	typedURL := "https://registry.npmjs.org/@types/" + name
	untypedURL := "https://registry.npmjs.org/" + name

	typedResponse, _ := http.Get(typedURL)
	untypedResponse, _ := http.Get(untypedURL)

	if api.Ok(untypedResponse.StatusCode) {
		err := exec.Command("yarn", "add", name).Run()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Cannot find package for " + name)
	}
	if api.Ok(typedResponse.StatusCode) {
		err := exec.Command("yarn", "add", "@types/"+name, "--dev").Run()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Cannot find typed package for " + name + ". Check if the untyped package has typings")
	}
}
