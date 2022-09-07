package server

import (
	"fmt"
	"maxd935/board/server/handler"
	"maxd935/board/structs"
	"net/http"
)

func Init_router(board [10][10]int, shipeSlic []structs.Shipe) {
	http.HandleFunc("/board", handler.BoardHandler(board))
	fmt.Println("/board")
	http.HandleFunc("/boats", handler.BoatHandler(shipeSlic))
	fmt.Println("/boats")
	http.HandleFunc("/hit", handler.HitHandler())
	fmt.Println("/hit")
}
