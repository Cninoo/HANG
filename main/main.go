package main

import (
	"fmt"
	"functions"
	"strings"
)

const theTotalNumberOfAttempts = 10

func main() {
	words, err := functions.DisplaysWordsFromFile("words.txt")

	if err != nil {
		fmt.Println("Impossible to read the file:", err)
		return
	}

	randomWord := functions.SelectRandomWord(words)              // choice of a random word
	revealCount := 1                                             // number of revealed letters
	hiddenWord := functions.MaskTheWord(randomWord, revealCount) // initial hidden word

	fmt.Println("Word to find:", hiddenWord)

	numberOfRemainingAttempts := theTotalNumberOfAttempts // initializing the number of attempts remaining
	guessedLetters := make(map[rune]bool)                 // initializing of the guessed letters

	// beginning of the game

	fmt.Println("Welcome to the game of Hangman!")
	fmt.Println("The word you need to guess is composed of", len(randomWord), "letters.")
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
			functions.Jose(numberOfRemainingAttempts)
		}

		fmt.Println("Be careful! The remaining number of attempts is:", numberOfRemainingAttempts)
		fmt.Println("Word:", functions.DisplayWord(randomWord, guessedLetters))

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

