package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"hangman/Jose"
)

// Constante pour le nombre total de tentatives
const theTotalNumberOfAttempts = 10

// Fonction pour masquer une partie du mot
func maskTheWord(word string, revealCount int) string {
	result := ""
	for i := 0; i < len(word); i++ { // Parcourt la string
		if i < revealCount { // Révèle une lettre si i vaut 0
			result += string(word[i]) + " " // Affiche la lettre devinée
		} else {
			result += "_ " // Laisse la lettre cachée si jamais elle n'est pas devinée
		}
	}
	return result
}

// Fonction pour afficher les mots à partir d'un fichier
func displaysWordsFromFile(filename string) ([]string, error) {
	f, err := os.Open(filename) // Ouvre le fichier que l'on souhaite
	if err != nil {             // Si erreur dans l'ouverture du fichier
		return nil, err
	}
	defer f.Close() // Gère la fermeture du fichier

	var words []string             // Liste de mots du fichier
	scanner := bufio.NewScanner(f) // Lit le fichier ligne par ligne
	for scanner.Scan() {
		words = append(words, scanner.Text()) // Ajout des mots dans la liste ci-dessus
	}

	if err := scanner.Err(); err != nil { // Gestion d'erreurs
		return nil, err
	}

	return words, nil
}

// Fonction pour choisir un mot au hasard à partir d'une liste de mots
func selectRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())     // Assure que le mot soit bien choisi au hasard à chaque exécution
	index := rand.Intn(len(words))       // Choisit un mot aléatoire
	return strings.ToUpper(words[index]) // Met le mot en Maj
}

// Fonction pour afficher le mot caché avec les lettres devinées
func displayWord(word string, guessedLetters map[rune]bool) string {
	displayedWord := ""
	for _, char := range word { // Boucle parcourant le mot à deviner
		if guessedLetters[char] { // Si lettre deviné, la lettre sera affichée et contenue dans la string initialisée au départ
			displayedWord += string(char)
		} else {
			displayedWord += "_" // Si elle n'est pas devinée, la lettre reste cachée
		}
		displayedWord += " "
	}
	return displayedWord
}

// lit le fichier et dit s'il y'a une erreur
func main() {
	words, err := displaysWordsFromFile("words.txt")
	if err != nil {
		fmt.Println("Impossible de  lire le fichier de mots:", err)
		return
	}

	randomWord := selectRandomWord(words)              // Choix d'un mot au hasard
	revealCount := 1                                   // Nombre de lettres révélées initialement
	hiddenWord := maskTheWord(randomWord, revealCount) //Mot caché initial

	fmt.Println("Word to find:", hiddenWord)

	numberOfRemainingAttempts := theTotalNumberOfAttempts // Initialisation du nombre de tentatives restantes
	guessedLetters := make(map[rune]bool)                 // Initialisation des lettres devinées

	//début du jeu

	fmt.Println("Welcome to the game of Hangman!")
	fmt.Println("The word you need to guess is composed of", len(randomWord), "lettres.")
	fmt.Println("Be careful! The number of remaining attempts is:", numberOfRemainingAttempts)

	for numberOfRemainingAttempts > 0 {
		fmt.Print("Enter any letter: ")
		var guess string
		fmt.Scanln(&guess)

		guessedLetter := rune(strings.ToUpper(guess)[0])
		found := false
		for _, char := range randomWord {
			if char == guessedLetter {
				guessedLetters[guessedLetter] = true
				found = true
			}
		}

		if !found {
			fmt.Println("The letter", string(guessedLetter), "is not present in the word..")
			numberOfRemainingAttempts--
			Jose(numberOfRemainingAttempts)
		}

		fmt.Println("Be careful! The remaining number of attempts is:", numberOfRemainingAttempts)
		fmt.Println("Mot présent:", displayWord(randomWord, guessedLetters))

		wordFound := true
		for _, char := range randomWord {
			if !guessedLetters[char] {
				wordFound = false
				break
			}
		}

		// fin du jeu
		if wordFound {
			fmt.Println("Congratulations, you found the word!")
			return
		}
	}

	fmt.Println("Sorry, you have no more attempts. The word was:", randomWord)
}
