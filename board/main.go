package main

import (
	"fmt"
	"maxd935/board/boardGame"
	"maxd935/board/server"
)

func main() {
	fmt.Println("----- Board Game -----")
	boardGameTab, shipeSlic := boardGame.CreateBoardGame()
	fmt.Println("shipeSlic ", shipeSlic)
	fmt.Println("")
	server.Init_router(boardGameTab, shipeSlic)

	myChannel1 := make(chan bool)
	go boardGame.PlayBoardGame(boardGameTab, myChannel1)
	server.Init_server(myChannel1)
}
