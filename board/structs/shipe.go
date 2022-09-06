package structs

import "fmt"

type Shipe struct {
	coors       Coors
	orientation int
	size        int
	hits        []Coors
}

func (shipe *Shipe) CreateShipe(orientation int, size int) {
	shipe.orientation = orientation
	shipe.size = size
}

func (shipe *Shipe) InitShipe(coors Coors) {
	shipe.coors = coors
}

func (shipe *Shipe) GetCoors() (int, int) {
	return shipe.coors.x, shipe.coors.y
}

func (shipe *Shipe) GetOrientation() int {
	return shipe.orientation
}

func (shipe *Shipe) GetSize() int {
	return shipe.size
}

func (shipe *Shipe) GetHit() []Coors {
	return shipe.hits
}

func (shipe *Shipe) addHit(coors Coors) {
	shipe.hits = append(shipe.hits, coors)
}

func (shipe *Shipe) PromptShipe() {
	fmt.Printf("coors => x: %d y: %d \n", shipe.coors.x, shipe.coors.y)
	fmt.Printf("orienation: %d \n", shipe.orientation)
	fmt.Printf("size: %d \n", shipe.size)
	fmt.Printf("hits: %d \n", shipe.hits)
}
