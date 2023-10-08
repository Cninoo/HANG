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

// Fonction pour afficher les mots à partir d'un fichier
func displaysWordsFromFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// Fonction pour choisir un mot au hasard à partir d'une liste de mots
func selectRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(words))
	return strings.ToUpper(words[index])
}

// Fonction pour afficher le mot caché avec les lettres devinées
func displayWord(word string, guessedLetters map[rune]bool) string {
	displayedWord := ""
	for _, char := range word {
		if guessedLetters[char] {
			displayedWord += string(char)
		} else {
			displayedWord += "_"
		}
		displayedWord += " "
	}
	return displayedWord
}
