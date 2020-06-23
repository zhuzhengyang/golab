package player

import "fmt"

type Player struct {
	Name string
}

func (r *Player) GetName() {
	r.setName("alice")
	fmt.Println("before patch ", r.Name)
}

func (r *Player) setName(n string) error {
	r.Name = n
	return nil
}
