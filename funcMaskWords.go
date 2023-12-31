package functions

func MaskTheWord(word string, revealCount int) string {
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