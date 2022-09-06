package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Valise Play")
	text := "Dans ma valise, il y a... "
	var slice []string
	max := 5
	for i := 0; i < max; i++ {
		fmt.Println(text)
		fmt.Println("Tour n°", i+1)
		var word string
		fmt.Print("Enter word: ")
		fmt.Scanf("%s", &word)
		slice = append(slice, word)
		if i == 0 {
			text = text + " " + word
		} else if i == max-1 {
			text = text + " et " + word
		} else {
			text = text + ", " + word
		}
		fmt.Println("Dans ma valise, il y a...", slice)
		fmt.Println(text)
	}
}

// Autre methode pour Rentrer une valeur
func reader() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
