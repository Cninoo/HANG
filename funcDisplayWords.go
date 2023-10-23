package functions

func DisplayWord(word string, guessedLetters map[rune]bool) string {
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
