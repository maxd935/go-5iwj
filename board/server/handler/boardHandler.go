package handler

import (
	"fmt"
	"net/http"
)

func BoardHandler(board [10][10]int) func(w http.ResponseWriter, req *http.Request) {
	boardHandlerReturn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(board)
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GOOD\n")
			fmt.Fprintf(w, "Size %d", len(board))
		}
	}
	return boardHandlerReturn
}
