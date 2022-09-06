package utils

import (
	"math/rand"
	"time"
)

func Random_number(min, max int) int {
	// fmt.Println(">> Random_number")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random_number := r1.Intn(max-min) + min
	// fmt.Println("random_number is ", random_number)
	return random_number
}
