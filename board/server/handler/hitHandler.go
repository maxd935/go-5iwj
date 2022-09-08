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
			fmt.Println("POST /add")
			// w.Header().Set("content-type", "application/json")
			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(boatResp)
		}
	}
	return boardHandlerReturn
}
