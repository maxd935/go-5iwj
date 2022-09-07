package handler

import (
	"fmt"
	"net/http"
)

func HitHandler() func(w http.ResponseWriter, req *http.Request) {
	boardHandlerReturn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hitHandler")
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "hitHandler\n")
		}
	}
	return boardHandlerReturn
}
