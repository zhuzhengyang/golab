package example

import (
	"github.com/zhuzhengyang/golab/hotfix/conf"
	"path/filepath"
	"testing"

	"github.com/zhuzhengyang/golab/hotfix/example/player"
	"github.com/zhuzhengyang/golab/hotfix/plugin"
)

func TestHotfix(t *testing.T) {
	p := plugin.NewPlugin()
	soFilePath := filepath.Join(conf.PluginPath, "patch.so")
	err := p.Build(soFilePath, &plugin.Config{
		Name:       "patch",
		ImportPath: "github.com/zhuzhengyang/golab/hotfix/example/patch",
		NewFunc:    "Patch",
	})

	if err != nil {
		t.Fatal(err)
	}

	pp := &player.Player{}
	pp.World()
	c, err := plugin.Load(soFilePath)
	if err != nil {
		t.Fatal(err)
	}
	err = plugin.Init(c)
	if err != nil {
		t.Fatal(err)
	}
	pp.Set("test1")
	pp.World()
}
