package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"./console"
)

type projectJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Version     string   `json:"version"`
	Main        string   `json:"main"`
	License     string   `json:"license"`
	Keywords    []string `json:"keywords,omitempty"`
	Author      string   `json:"author,omitempty"`
}

func buildPackage() error {
	path, err := os.Getwd()
	if err != nil {return err}

	var packageJSON projectJSON

	basename := filepath.Base(path)
	packageJSON.Name = console.Prompt("package name ("+basename+"): ", basename)
	packageJSON.Version = console.Prompt("version (1.0.0): ", "1.0.0")
	packageJSON.Description = console.Prompt("description: ", "")
	packageJSON.Main = console.Prompt("entry point (index.js): ", "index.js")
	packageJSON.License = console.Prompt("license (MIT): ", "MIT")
	packageJSON.Author = console.Prompt("author: ", "")
	keywords := console.Prompt("keywords: ", "")
	for _, k := range strings.Split(keywords, "\n") {
		if t := strings.TrimSpace(k); len(t) != 0 {
			packageJSON.Keywords = append(packageJSON.Keywords, t)
		}
	}

	file, err := os.Create("package.json")
	if err != nil {return err}

	data, err := json.MarshalIndent(packageJSON, "", "  ")
	if err != nil {return err}

	_, err = file.Write(data)
	if err != nil {return err}

	return nil
}
