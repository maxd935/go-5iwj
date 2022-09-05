package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Clock")
	date, time := get_datetime()
	fmt.Printf("Nous sommes le %s, il est %s.", date, time)
}

func get_datetime() (string, string) {
	fmt.Println("get_datetime()")
	datetime := time.Now()
	fmt.Println(datetime.Location())

	date := datetime.Format("1 January")
	time := datetime.Format("15h04")
	return date, time
}
