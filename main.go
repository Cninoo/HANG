package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const theTotalNumberOfAttempts = 10

func Jose(numberOfAttemps int) { //fonction qui affiche la position du pendu par rapport au nombre de vie restante
	var arrayJose []byte
	initNumberOfLife := 10
	pos := initNumberOfLife - numberOfAttemps
	content, err := os.ReadFile("hangman.txt")
	raid := 1
	fr := 1
	if err != nil {
		fmt.Println("[fichier non trouvé]")
	} else {
		for i := 0; i < len(content); i++ {
			if content[i] == 10 {
				raid++
			}
			if raid != 8 {
				arrayJose = append(arrayJose, content[i])
			}
			if raid == 8 && pos != fr {
				arrayJose = []byte{}
				raid = 1
				fr++
			}
			if raid == 8 && pos == fr {
				jose := (string(arrayJose))
				fmt.Println(jose)
				break
			}
		}
	}
}

func maskTheWord(word string, revealCount int) string {
	result := ""
	for i := 0; i < len(word); i++ { // browse the length of the word
		if i < revealCount { // if i is lower than revealCount he must print a caracter
			result += string(word[i]) + " " // print the guessed letters
		} else {
			result += "_" // print an underscore if the letter isn't guess
		}
	}
	return result
}

func displaysWordsFromFile(filename string) ([]string, error) {
	f, err := os.Open(filename) // open the file
	if err != nil {             // error in the opening of the file
		return nil, err
	}
	defer f.Close() // ensure the file close correctly

	var words []string             // variable in wich the words of the file are stored
	scanner := bufio.NewScanner(f) // read the file line by line
	for scanner.Scan() {
		words = append(words, scanner.Text()) // add the words of the file to the word list
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func selectRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())     // ensure that the selection will be random
	index := rand.Intn(len(words))       // extract a random word from the index
	return strings.ToUpper(words[index]) // convert the word in uppercase
}

func displayWord(word string, guessedLetters map[rune]bool) string {
	displayedWord := ""
	for _, char := range word { // browse the word to guess
		if guessedLetters[char] {
			displayedWord += string(char) // the guessed letter is add to the word
		} else {
			displayedWord += "_" // if the letter is not guess, it print an underscore
		}
		displayedWord += " " // add a space after each caracter
	}
	return displayedWord
}

func main() {
	words, err := displaysWordsFromFile("words.txt")
	if err != nil {
		fmt.Println("Impossible to read the file:", err)
		return
	}

	randomWord := selectRandomWord(words)              // choice of a random word
	revealCount := 1                                   // number of revealed letters
	hiddenWord := maskTheWord(randomWord, revealCount) // initial hidden word

	fmt.Println("Word to find:", hiddenWord)

	numberOfRemainingAttempts := theTotalNumberOfAttempts // initializing the number of attempts remaining
	guessedLetters := make(map[rune]bool)                 // initializing of the guessed letters

	// beginning of the game

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
		}

		fmt.Println("Be careful! The remaining number of attempts is:", numberOfRemainingAttempts)
		fmt.Println("Word:", displayWord(randomWord, guessedLetters))

		wordFound := true
		for _, char := range randomWord {
			if !guessedLetters[char] {
				wordFound = false
				break
			}
		}

		// end of the game
		if wordFound {
			fmt.Println("Congratulations, you found the word!")
			return
		}
	}

	fmt.Println("Sorry, you have no more attempts. The word was:", randomWord)
}
