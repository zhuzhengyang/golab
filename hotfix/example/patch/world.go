package patch

import (
	"fmt"
	"reflect"
	_ "unsafe"

	"github.com/zhuzhengyang/golab/hotfix/example/player"

	"github.com/agiledragon/gomonkey"
)

//go:linkname  setName github.com/zhuzhengyang/golab/hotfix/example/player.(*Player).setName
func setName(r *player.Player, n string) error

func FixWorld(r *player.Player) {
	setName(r, "patch ok")
	fmt.Println("Wonderful World ", r.Name)
}

func Patch() {
	var d *player.Player
	fmt.Println("patch exec")
	gomonkey.ApplyMethod(reflect.TypeOf(d), "World", FixWorld)
}
