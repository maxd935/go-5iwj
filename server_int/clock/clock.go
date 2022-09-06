package clock

import (
	"fmt"
	"net/http"
	"time"
)

func ClockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClockHandler")
	_, time := get_datetime()
	fmt.Fprintf(w, ">%s", time)
}

func get_datetime() (string, string) {
	datetime := time.Now()

	date := datetime.Format("1 January")
	time := datetime.Format("15h04")
	return date, time
}
