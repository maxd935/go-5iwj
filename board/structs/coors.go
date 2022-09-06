package structs

import "fmt"

type Coors struct {
	x, y int
}

func (coors *Coors) InitCoors(x, y int) {
	coors.x = x
	coors.y = y
}

func (coors *Coors) PromptCoors() {
	fmt.Printf("x: %d y: %d \n", coors.x, coors.y)
}

func (coors *Coors) GetCoors() (x, y int) {
	return coors.x, coors.y
}
