package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Charge le contenu du fichier texte dans une chaîne de caractères
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Divise le contenu en mots en utilisant l'espace comme délimiteur
	words := strings.Fields(string(content))

	// Initialise le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Choisis un mot au hasard
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]

	// Affiche le mot au hasard
	fmt.Println(randomWord)
}
