// Package game sets up a new game
package game

import (
	"bufio"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"fmt"
	"log"
	"os"
)

//Player defines a player
type Player struct {
    Name string;
    Guess uint64;
}

//Gets returns player Name
func getName(NameReader *bufio.Reader) string{
    fmt.Print("What should I call you? ")
    playerName, err := NameReader.ReadString('\n')
    if err != nil {
        log.Fatal("Could not get player Name: ", err)
    }

    playerName = strings.TrimSuffix(playerName, "\n")

    return playerName 
}
//getGuess returns the players guess
func getGuess(guessReader *bufio.Reader, player string) uint64 {
    fmt.Printf("Take your guess %s: ", player)
    guess, err := guessReader.ReadString('\n')

    //remove new line
    guess = strings.TrimSuffix(guess, "\n")
    if err != nil {
        log.Fatal("Error reading input:", err)
    }
    //convert to an int
    playerGuess, err := strconv.Atoi(guess)
    if err != nil {
        log.Fatal("Could not convert to int:", err)
    }

    return uint64(playerGuess)
}

//getInput returns playerName and playerGuess
func getInput() (uint64, string) {
    reader := bufio.NewReader(os.Stdin)
    playerName := getName(reader)
    playerGuess := getGuess(reader, playerName)

    return playerGuess, playerName
}

//generateNumber returns the game answer value
func generateNumber() uint64 {
    rand.Seed(time.Now().Unix())
    gameValue := rand.Intn((50 - 1) + 1)

    return uint64(gameValue)
}

//Starts the game
func StartGame()  {
    println("\nThe billion dollar guessing game\nAllowed guesses [1-50]\nYou only get 3 chances\n")
    player := &Player{}
    input, playerName := getInput()
    player.Name = playerName
    player.Guess= input
    correctAnswer := generateNumber()
    if player.Guess == correctAnswer {
        fmt.Printf("\nIts a match\nYour Guess {%d} = correctAnswer{%d}", player.Guess, correctAnswer)
    } else {
        fmt.Printf("\nThe correctAnswer was {%d}\n", correctAnswer)
        fmt.Println("Lets go again")
    } 
}
