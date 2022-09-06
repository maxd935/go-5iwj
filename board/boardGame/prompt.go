package boardGame

import "fmt"

func PromptBoardGame(boardGame [10][10]int) {
	fmt.Println("> PromptBoardGame")
	for i, _ := range boardGame {
		fmt.Println("---------------------")
		for j, _ := range boardGame[i] {
			fmt.Print("|")
			fmt.Print(boardGame[i][j])
		}
		fmt.Println("|")
	}
}
