package hotfix

import (
	"testing"

	"hotfix/player"
	"hotfix/plugin"
)

func TestHotfix(t *testing.T) {
	p := plugin.NewPlugin()
	err := p.Build("/tmp/patch.so", &plugin.Config{
		Name:       "patch",
		ImportPath: "hotfix/patch",
		NewFunc:    "Patch",
	})

	if err != nil {
		t.Fatal(err)
	}

	pp := &player.Player{}
	pp.World()
	c, err := plugin.Load("/tmp/patch.so")
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
