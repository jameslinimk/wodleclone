package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

/* ---------------------------------- Words --------------------------------- */
//go:embed words.txt
var wordFile string
var words = strings.Split(wordFile, ",")

//go:embed wordsGuesses.txt
var wordGuessesFile string
var wordGuesses = strings.Split(wordGuessesFile, ",")

/* ------------------------------- Game stuff ------------------------------- */
var word string // Word trying to guess
var maxGuessCount = 6
var guesses = strings.Split(strings.Repeat("e", maxGuessCount), "")

/* ---------------------------------- Fonts --------------------------------- */
var fontDeafult = color.New(color.FgWhite)
var fontSuccess = color.New(color.FgGreen)
var fontError = color.New(color.FgRed)
var fontBlankBox = color.New(color.BgWhite).Add(color.FgBlack)
var fontIncludesBox = color.New(color.BgYellow).Add(color.FgBlack)
var fontCorrectBox = color.New(color.BgGreen).Add(color.FgBlack)

func isWord(word string) bool {
	found := false
	fmt.Println(wordGuesses)
	for _, w := range wordGuesses {
		if w == word {
			found = true
			break
		}
	}
	return found
}

func isNewGuess(word string) bool {
	new := true
	for _, w := range guesses {
		if w == word {
			new = false
			break
		}
	}
	return new
}

func getWordInput() string {
	var nextGuess int
	for i, v := range guesses {
		if v == "e" {
			nextGuess = i
			break
		}
	}

	fontDeafult.Printf("(%d/%d) Guess:\n", nextGuess, maxGuessCount)

	var input string
	for {
		fmt.Scanln(&input)

		if isWord(strings.ToLower(input)) {
			if isNewGuess(strings.ToLower(input)) {
				break
			} else {
				fontError.Println("Guess a new word!")
			}
		} else {
			fontError.Println("Not a valid word")
		}
	}

	return input
}

func drawUI() {
	// fontDeafult.Printf("Word: %s\n", word)

	for i := 0; i < maxGuessCount; i++ {
		for x := 0; x < 5; x++ {
			if guesses[i] != "e" {
				letter := string(guesses[i][x])
				index := strings.Index(word, letter)

				if index == -1 {
					// Character is not in word
					fontBlankBox.Printf(" %s ", letter)
				} else if string(word[x]) == letter {
					// Character is in the right place
					fontCorrectBox.Printf(" %s ", letter)
				} else {
					// Character is in the word but not in the right place

					// Before, check if it already has been checked by the correct place
					found := false
					for i, v := range strings.Split(guesses[i], "") {
						if string(word[i]) == v {
							found = true
							break
						}
					}
					if found {
						fontBlankBox.Printf(" %s ", letter)
					} else {
						fontIncludesBox.Printf(" %s ", letter)
					}
				}
			} else {
				fontBlankBox.Printf("   ")
			}

			fontDeafult.Print(" ")
		}
		fontDeafult.Println()
		fontDeafult.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	word = words[rand.Intn(len(words))] // Word trying to guess

	for {
		drawUI()

		guess := getWordInput()

		var nextGuess int
		for i, v := range guesses {
			if v == "e" {
				nextGuess = i
				break
			}
		}

		guesses[nextGuess] = guess

		if guess == word {
			drawUI()
			fontSuccess.Printf("\n🎉 You won in %d tries!\n", nextGuess+1)
			break
		}

		if nextGuess == maxGuessCount-1 {
			drawUI()
			fontError.Printf("\n😢 You lose, the word was %s\n", word)
			break
		}
	}

	fmt.Println("Press any key to exit")
	fmt.Scanln()
}
