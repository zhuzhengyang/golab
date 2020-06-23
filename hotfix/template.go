package hotfix

var (
	tmpl = `package main

import (
	"github.com/zhuzhengyang/golab/hotfix"
	"{{.ImportPath}}"
)
var Plugin = hotfix.Config{
	ImportPath: "{{.ImportPath}}",
	NewFunc: {{.NewFunc}},
}
`
)
