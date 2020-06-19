package plugin

var (
	tmpl = `package main

import (
	"hotfix/plugin"
	"{{.ImportPath}}"
)
var Plugin = plugin.Config{
	Name: "{{.Name}}",
	ImportPath: "{{.ImportPath}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
