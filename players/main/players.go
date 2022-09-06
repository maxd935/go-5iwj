package main

import (
	"fmt"
	"math/rand"
	"time"

	"maxd935/players/fileP"
)

type player struct {
	nom    string
	prenom string
	pseudo string
	score  int
	PdV    int
}

func (playerOne *player) addPlayer() {
	nom, prenom, pseudo := getInfo()
	playerOne.nom = nom
	playerOne.prenom = prenom
	playerOne.pseudo = pseudo
}

func main() {
	var players []player
	var choix string
	for {
		if len(players) > 0 {
			fmt.Println("----- MENU -----")
			fmt.Println("1 - Play")
			fmt.Println("2 - Ajouter un Joueur")
			fmt.Println("3 - Afficher les Joueurs")
			fmt.Println("0 - Quitter")
			fmt.Print("Entrer le numero de votre choix: ")
			fmt.Scanf("%s", &choix)
			fmt.Println("")
			switch choix {
			case "1":
				fmt.Println("Play !!")
				Play(players)
				fileP.ReadFile()
			case "2":
				fmt.Println("New Player !!")
				var playerNew player
				playerNew.addPlayer()
				players = append(players, playerNew)
				getPlayers(players)
			case "3":
				fmt.Println("Afficher les Joueurs")
				getPlayers(players)
			case "0":
				fmt.Println("Bye !")
				return
			default:
				fmt.Println("Mauvais choix !")
			}
		} else {
			fmt.Println("New Player !")
			var playerNew player
			playerNew.addPlayer()
			players = append(players, playerNew)
			getPlayers(players)
		}
	}
}

func Play(players []player) {
	fmt.Println("")
	fmt.Println("Trouve le bon chiffre")
	max := 100
	nbrPartieMax := 5
	for i := 0; i < nbrPartieMax; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		number_find := r1.Intn(max)
		fmt.Println("")
		getPlayers(players)
		fmt.Println("Partie n°", i+1)
		nbrRound := 0
		stopGame := false
		for !stopGame {
			nbrRound++
			fmt.Println("\tRound n°", nbrRound)
			for j, player := range players {
				fmt.Println("Player:", player.pseudo)
				stopGame = getNumber(number_find, &players[j])
				if stopGame {
					break
				}
			}
		}
	}
}

func getNumber(good_number int, player *player) bool {
	var number int
	// fmt.Println("Player:", player)

	fmt.Print("Entrer un chiffre: ")
	fmt.Scanf("%d", &number)

	if number < good_number {
		fmt.Println("\tTrop petit")
		player.score++
		return false
	} else if number > good_number {
		fmt.Println("\tTrop grand")
		player.score++
		return false
	} else {
		fmt.Println("\tT'es trop fort ! Le chiffer était ", good_number)
		player.score--
		return true
	}
}

func getPlayers(players []player) {
	fmt.Println("Il y a ", len(players), " joueurs.")
	for i, player := range players {
		fmt.Println("Player n°", i+1)
		fmt.Printf("\tNom : %v Prenom : %v Pseudo: %v Score: %d PdV: %d \n\n", player.nom, player.prenom, player.pseudo, player.score, player.PdV)
	}
}

func getInfo() (string, string, string) {
	fmt.Println("----- Create Player ------")

	var nom string
	fmt.Print("Entrer un nom: ")
	fmt.Scanf("%s", &nom)

	var prenom string
	fmt.Print("Entrer un prenom: ")
	fmt.Scanf("%s", &prenom)

	var pseudo string
	fmt.Print("Entrer un pseudo: ")
	fmt.Scanf("%s", &pseudo)
	return nom, prenom, pseudo
}
