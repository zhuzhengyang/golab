package patch

import (
	"fmt"
	"reflect"
	_ "unsafe"

	"github.com/zhuzhengyang/golab/hotfix"

	"github.com/zhuzhengyang/golab/hotfix/example/player"

	"github.com/agiledragon/gomonkey"
)

var PluginName = "example_patch"

//go:linkname  setName github.com/zhuzhengyang/golab/hotfix/example/player.(*Player).setName
func setName(r *player.Player, n string) error

func FixGetName(r *player.Player) {
	setName(r, "bob")
	fmt.Println("after patch ", r.Name)
}

func PatchPlayerWorld() {
	var d *player.Player
	fmt.Println("patch exec")
	patch := gomonkey.ApplyMethod(reflect.TypeOf(d), "GetName", FixGetName)
	hotfix.RegisterPatch(PluginName, patch)
}
