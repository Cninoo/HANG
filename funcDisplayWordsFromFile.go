package functions

import (
	"bufio"
	"os"
)

func DisplaysWordsFromFile(filename string) ([]string, error) {
	f, err := os.Open(filename) // Ouvre le fichier que l'on souhaite
	if err != nil {             // Si erreur dans l'ouverture du fichier
		return nil, err
	}
	defer f.Close()                // Gère la fermeture du fichier
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
