package functions

import (
	"fmt"
	"os"
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
