package ascii

import (
	"fmt"
	"strings"
)

func ProcessWords(words string) {
	wordsArr := strings.Split(words, "\\n")

	for _, word := range wordsArr {
		if word == "" {
			fmt.Println()
			continue
		}
		PrintChar(word)
	}
}
