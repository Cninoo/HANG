package main

// Constante pour le nombre total de tentatives
const theTotalNumberOfAttempts = 10

// Fonction pour masquer une partie du mot
func maskTheWord(word string, revealCount int) string {
	result := ""
	for i := 0; i < len(word); i++ {
		if i < revealCount {
			result += string(word[i]) + " "
		} else {
			result += "_ "
		}
	}
	return result
}
 