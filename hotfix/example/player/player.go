package player

import "fmt"

// Test Struct
type Player struct {
	Name string
}

// Print Name
func (r *Player) GetName() {
	r.setName("alice")
	fmt.Println("before patch ", r.Name)
}

func (r *Player) setName(n string) error {
	r.Name = n
	return nil
}
