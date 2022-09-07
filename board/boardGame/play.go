package boardGame

import "fmt"

func PlayBoardGame(myBoard [10][10]int, myChannel1 chan bool) {
	res := <-myChannel1
	if res {
		fmt.Println("> PlayBoardGame")
		var pseudo string
		fmt.Print("Entrer votre pseudo: ")
		fmt.Scanf("%s", &pseudo)
		fmt.Println("Hey, ", pseudo)
		actions(myBoard)
	}
}

func actions(board [10][10]int) {
	var choix string
	// if len(boats) > 0 {
	if true {
		for {
			fmt.Println("----- MENU -----")
			fmt.Println("1 - Ajouter adversaire")
			fmt.Println("2 - Etat du Plateau")
			fmt.Println("3 - Nombre de bateau actuel")
			fmt.Println("4 - Lancer une attaque")
			fmt.Println("0 - Quitter")
			fmt.Print("Entrer le numero de votre choix: ")
			fmt.Scanf("%s", &choix)
			fmt.Println("")
			switch choix {
			case "1":
				fmt.Println("New adversaire !!")
				addAdversaire()
			case "2":
				fmt.Println("Etat du Plateau")
				etatPlateau()
			case "3":
				fmt.Println("Nombre de bateau actuel")
				nbrBateau()
			case "4":
				fmt.Println("Lancer une attaque")
				launchAttack()
			case "0":
				fmt.Println("Bye !")
				return
			default:
				fmt.Println("Mauvais choix !")
			}
		}
	} else {
		fmt.Println("Fin de Partie")
		return
	}
}
