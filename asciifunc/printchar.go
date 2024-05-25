package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintChar(word string) {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error reading from file")
		return
	}
	filecontent := strings.Split(string(file), "\n")

	if strings.Contains(string(file), "o") {
		filecontent = strings.Split(string(file), "\r\n")
	}

	for l := 1; l <= 8; l++ {
		for _, char := range word {
			index := int(char-32) * 9
			fmt.Print(filecontent[index+l])
		}
		fmt.Println()
	}
}
