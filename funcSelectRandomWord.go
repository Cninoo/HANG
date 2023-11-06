package functions

import (
	"math/rand"
	"strings"
	"time"
)

func SelectRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())     // ensure that the selection will be random
	index := rand.Intn(len(words))       // extract a random word from the index
	return strings.ToUpper(words[index]) // convert the word in uppercase
}
