package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func hideWord(word string, revealCount int) string {
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

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var words []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(words) > 0 {
		index := rand.Intn(len(words))
		chosenWord := words[index]
		revealCount := 1
		hiddenWord := hideWord(chosenWord, revealCount)

		fmt.Println("revealCount:", revealCount)
		fmt.Println("hiddenWord:", hiddenWord)
	} else {
		fmt.Println("No words found in the file.")
	}


