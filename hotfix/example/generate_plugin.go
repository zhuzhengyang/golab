package main

import (
	"log"

	"github.com/zhuzhengyang/golab/hotfix"

	"github.com/zhuzhengyang/golab/hotfix/example/patch"
)

func main() {
	hotfix.SetPluginPath("./")
	err := hotfix.Build(patch.PluginName, &hotfix.Config{
		ImportPath: "github.com/zhuzhengyang/golab/hotfix/example/patch",
		NewFunc:    "patch.PatchPlayerWorld",
	})

	if err != nil {
		log.Fatal(err)
	}
}
