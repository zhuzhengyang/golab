package main

import (
	"time"

	"github.com/zhuzhengyang/golab/hotfix"
	"github.com/zhuzhengyang/golab/hotfix/example/player"
)

func main() {
	hotfix.Watch("")
	defer hotfix.StopWatch()

	pp := &player.Player{}
	c := time.Tick(1 * time.Second)
	for _ = range c {
		pp.GetName()
	}
}
