package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Rool Play")
	res1 := random_int(6)
	res2 := random_int(20)
	res3 := random_int(100)
	fmt.Printf("|%02d|  |%02d|  |%03d|", res1, res2, res3)
}

func random_int(size int) int {
	fmt.Println("random_int() with", size)
	randomInt := rand.Intn(size)
	return randomInt
}
