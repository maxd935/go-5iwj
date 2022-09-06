package boardGame

import (
	"fmt"
	"maxd935/board/structs"
	"maxd935/board/utils"
)

const SIZE_BOARD = 10

const MIN_SIZE_SHIPE = 1
const MAX_SIZE_SHIPE = 6
const MAX_SHIPE = 5

const ORIENTATION_ORIZONTALE = 0
const ORIENTATION_VERTICALE = 2

func CreateBoardGame() ([10][10]int, []structs.Shipe) {
	fmt.Println("> CreateBoardGame")
	boardGame := createTab()
	var shipeSlic []structs.Shipe
	for i := 0; i < MAX_SHIPE; i++ {
		shipeNew := createShipe()
		addShipe(&shipeNew, &boardGame, i+1)
		shipeSlic = append(shipeSlic, shipeNew)
	}
	return boardGame, shipeSlic
}

func createTab() [10][10]int {
	fmt.Println(">> createTab")
	var boardGame [SIZE_BOARD][SIZE_BOARD]int
	return boardGame
}

func createShipe() structs.Shipe {
	fmt.Println(">> createShipe")
	new_orientation := utils.Random_number(ORIENTATION_ORIZONTALE, ORIENTATION_VERTICALE)
	new_size := utils.Random_number(MIN_SIZE_SHIPE, MAX_SIZE_SHIPE)

	var shipeNew structs.Shipe
	shipeNew.CreateShipe(new_orientation, new_size)
	// shipeNew.PromptShipe()
	return shipeNew
}

func createCoors() structs.Coors {
	fmt.Println(">> createCoors")
	new_coors_x := utils.Random_number(0, SIZE_BOARD)
	new_coors_y := utils.Random_number(0, SIZE_BOARD)

	var new_coors structs.Coors
	new_coors.InitCoors(new_coors_x, new_coors_y)
	// new_coors.PromptCoors()

	return new_coors
}

func isValidCoorsBoard(coors structs.Coors, shipe structs.Shipe, boardGame [10][10]int) bool {
	fmt.Println(">> isValidCoorsBoard")
	x, y := coors.GetCoors()

	if shipe.GetOrientation() == 1 {
		if len(boardGame) > shipe.GetSize()+x {
			for i := 0; i < shipe.GetSize(); i++ {
				if boardGame[x+i][y] != 0 {
					fmt.Println("FALSE : boat present")
					return false
				}
			}
		} else {
			fmt.Println("FALSE : limit max tab X")
			return false
		}

	} else {
		if len(boardGame) > shipe.GetSize()+y {
			for i := 0; i < shipe.GetSize(); i++ {
				if boardGame[x][y+i] != 0 {
					fmt.Println("FALSE : boat present Y")
					return false
				}
			}
		} else {
			fmt.Println("FALSE : limit max tab")
			return false
		}
	}
	fmt.Println("TRUE")
	return true
}

func addShipe(shipeNew *structs.Shipe, boardGame *[10][10]int, num_shipe int) {
	for {
		coors := createCoors()
		if isValidCoorsBoard(coors, *shipeNew, *boardGame) {
			shipeNew.InitShipe(coors)
			x, y := shipeNew.GetCoors()
			if shipeNew.GetOrientation() == 1 {
				for i := 0; i < shipeNew.GetSize(); i++ {
					boardGame[x+i][y] = num_shipe
				}

			} else {
				for i := 0; i < shipeNew.GetSize(); i++ {
					boardGame[x][y+i] = num_shipe
				}
			}
			shipeNew.PromptShipe()
			return
		}
	}
}
