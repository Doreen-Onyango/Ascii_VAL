package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// readAscii reads ASCII art from the file and returns a map of runes to art strings.
func readAscii(filename string) (map[rune]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	asciiArtMap := make(map[rune]string)
	scanner := bufio.NewScanner(file)

	var currentRune rune = 32 // Start from space character
	var artLines []string
	var lineCount int
	var maxLen int

	for scanner.Scan() {
		line := scanner.Text()
		artLines = append(artLines, line)
		lineCount++

		// Find the length of the line up to the last character plus one space
		lastCharIndex := strings.LastIndexFunc(line, func(c rune) bool {
			return !unicode.IsSpace(c)
		})
		if lastCharIndex+1 > maxLen {
			maxLen = lastCharIndex + 1
		}

		// Every 9 lines, map the ASCII art to the current rune and move to the next rune
		if lineCount%9 == 0 {
			for _, artLine := range artLines[1:] {
				// If the line is shorter than maxLen, append spaces
				if len(artLine) < maxLen {
					artLine += strings.Repeat(" ", maxLen-len(artLine))
				}
				asciiArtMap[currentRune] += artLine[:maxLen] + "\n"
			}
			artLines = artLines[:0] // Reset artLines
			currentRune++           // Move to the next rune
			maxLen = 0              // Reset maxLen
			if currentRune > 126 {  // Stop if we've reached the tilde (~)
				break
			}
		}
	}

	return asciiArtMap, scanner.Err()
}

// generateAsciiArt generates combined ASCII art from input string.
func GenerateAsciiArt(input string, asciiArtMap map[rune]string) string {
	// Split the input into runes
	runes := []rune(input)

	// Create a slice to hold the ASCII art lines for each rune
	var asciiArtLines [][]string

	// For each rune in the input, get its ASCII art and split it into lines
	for _, r := range runes {
		if art, ok := asciiArtMap[r]; ok {
			lines := strings.Split(art, "\n")
			asciiArtLines = append(asciiArtLines, lines)
		}
	}

	// Concatenate the ASCII art lines horizontally
	var result string
	for i := 0; i < len(asciiArtLines[0]); i++ {
		for _, lines := range asciiArtLines {
			if i < len(lines) {
				result += lines[i]
			}
		}
		result += "\n"
	}

	return strings.TrimSuffix(result, "\n")
}

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input string>")
		os.Exit(1)
	}

	// Read the ASCII art from the file
	asciiArtMap, err := readAscii("standard.txt")
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		os.Exit(1)
	}

	// Process each argument
	for _, arg := range os.Args[1:] {
		result := strings.Split(arg, "\\n")

		for _, word := range result {
			if word == "" {
				fmt.Println()
			} else {
				art := GenerateAsciiArt(word, asciiArtMap)
				fmt.Print(art)
			}
		}
	}
}
