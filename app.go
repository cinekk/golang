package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
	"strconv"
)

func intro() {
	fmt.Println("ELO bijacz. Welkum to da game!")
	fmt.Println("System will generate a random number in rage 1 - 100")
	fmt.Println("You need to guess dat.")
	prompt()
}

func prompt() {
	fmt.Println("")
	fmt.Println("Type in yer guess:")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Intn(100)

	intro()

	for scanner.Scan() {
		
		guess := scanner.Text()

		var guessed int

		guessed, err := strconv.Atoi(guess)

		if err != nil {
			fmt.Println("NO INPUT BIJACZ")
		}

		var msg string

		if guessed != randNum {
			if guessed > randNum {
				msg = "too high"
			} else {
				msg = "too low"
			}
			fmt.Println("Guess is no good! Guesss", msg, "!!!")

			prompt()
		} else {
			fmt.Println("GUESSS OKAYYYY")
			os.Exit(0)
		}
	}

	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }	
}