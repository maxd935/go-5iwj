package fileP

import (
	"fmt"
	"os"
)

func ReadFile() {
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
