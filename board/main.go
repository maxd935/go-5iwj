package main

import (
	"fmt"
	"maxd935/board/boardGame"
)

func main() {
	fmt.Println("----- Board Game -----")
	boardGameTab, shipeSlic := boardGame.CreateBoardGame()
	fmt.Println("shipeSlic ", shipeSlic)
	boardGame.PromptBoardGame(boardGameTab)
}
