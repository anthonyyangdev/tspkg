package registry

import (
	"fmt"
	"net/http"
	"os/exec"
)

func ok(code int) bool {
	return 200 <= code && code < 300
}

func Get(name string) {
	typedURL := "https://registry.npmjs.org/@types/" + name
	untypedURL := "https://registry.npmjs.org/" + name

	typedResponse, _ := http.Get(typedURL)
	untypedResponse, _ := http.Get(untypedURL)

	if ok(untypedResponse.StatusCode) {
		err := exec.Command("yarn", "add", name).Run()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Cannot find package for " + name)
	}
	if ok(typedResponse.StatusCode) {
		err := exec.Command("yarn", "add", "@types/"+name, "--dev").Run()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Cannot find typed package for " + name + ". Check if the untyped package has typings")
	}
}

// export async function tsGet(name: string) {
// 	const response = await fetch(typedGetUrl(name));
// 	const processOptions = {cwd: process.cwd()};
// 	if (response.ok) {
// 	  const response = await fetch(regularUrl(name));
// 	  if (response.ok) {
// 		child_process.exec(`yarn add ${name}`, processOptions);
// 		child_process.exec(`yarn add @types/${name} --dev`, processOptions);
// 		return "Installed packages";
// 	  } else {
// 		return Promise.reject(`Could not find Node package: ${name}`);
// 	  }
// 	} else {
// 	  const response = await fetch(regularUrl(name));
// 	  if (!response.ok) {
// 		return Promise.reject(`Package ${name} not found`);
// 	  } else {
// 		child_process.exec(`yarn add ${name}`, processOptions);
// 		return "Installed possibly only untyped package. Check if the package already includes types.";
// 	  }
// 	}
//   }
