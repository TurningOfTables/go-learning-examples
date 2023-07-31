// A quick version of a PIN code game
// A random number is generated which the user must guess
// After each guess the user is told whether the number is a) Correct, b) Present but in a different position or c) Incorrect

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var pinLength int = 5
var valid = regexp.MustCompile(`^[0-9]*$`)

const (
	RIGHT    = "✔"
	WRONG    = "✘"
	POSITION = "?"
)

func main() {

	printTitle()
	readyPrompt()
	pin := generatePin()
	doGuesses(pin)

}

func generatePin() []int {
	var pin []int

	for i := 0; i < pinLength; i++ {
		pin = append(pin, rand.Intn(9))
	}

	fmt.Printf("%v number PIN generated\n", pinLength)
	return pin
}

func doGuesses(pin []int) {
	guess := guessPrompt()

	if guess == nil {
		fmt.Printf("You're giving up! The PIN was: %v", pin)
		return
	}

	isCorrect, details := checkGuess(pin, guess)

	if !isCorrect {

		fmt.Println("Incorrect PIN :(")
		for k, v := range guess {
			fmt.Printf("%v - %v\n", v, details[k])
		}
		doGuesses(pin)
	} else {
		fmt.Println("Correct!")
	}
}

func checkGuess(pin []int, guess []int) (bool, []string) {
	guessDetails := make([]string, pinLength)
	isCorrect := false

	for d := range guessDetails {
		guessDetails[d] = RIGHT
	}

	if pinMatches(pin, guess) {
		isCorrect = true
		for d := range guessDetails {
			guessDetails[d] = RIGHT
		}
		return isCorrect, guessDetails
	}

	for k := range pin {
		if pin[k] == guess[k] {
			guessDetails[k] = RIGHT
		} else if pinContains(pin, guess[k]) {
			guessDetails[k] = POSITION
		} else {
			guessDetails[k] = WRONG
		}
	}

	return isCorrect, guessDetails
}

func guessPrompt() []int {
	var g string

	fmt.Printf("Enter your %v number guess or (quit) to give up: ", pinLength)
	fmt.Scan(&g)

	if g == "quit" {
		return nil
	}

	validGuess := valid.MatchString(g)
	if !validGuess {
		fmt.Println("Oops - you can only enter numbers in your guess")
		guessPrompt()
	}

	if len(g) != pinLength {
		fmt.Printf("Oops - make sure your guess is the right length (%v)\n", pinLength)
		guessPrompt()
	}

	stringSlice := strings.Split(g, "")

	var intSlice []int
	for _, s := range stringSlice {
		stringAsInt, _ := strconv.Atoi(s)
		intSlice = append(intSlice, stringAsInt)
	}

	return intSlice
}

func readyPrompt() {
	var i string
	fmt.Print("Ready to start? (y): ")
	fmt.Scan(&i)

	if strings.ToLower(i) == "y" {
		return
	} else {
		readyPrompt()
	}
}

func printTitle() {
	fmt.Println(`
	========
	PIN CODE
	========
	`)
}

func pinContains(pin []int, i int) bool {
	for _, v := range pin {
		if v == i {
			return true
		}
	}

	return false
}

func pinMatches(pin []int, guess []int) bool {
	for k := range pin {
		if pin[k] != guess[k] {
			return false
		}
	}
	return true
}
