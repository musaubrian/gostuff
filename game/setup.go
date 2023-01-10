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
    name string;
    guess uint64;
}
//Gets returns player name
func getName(nameReader *bufio.Reader) string{
    fmt.Print("What should I call you?> ")
    playerName, err := nameReader.ReadString('\n')
    if err != nil {
        log.Fatal("Could not get player name: ", err)
    }

    playerName = strings.TrimSuffix(playerName, "\n")

    return playerName 
}

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
//getInput returns the players guess as an int
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
func Setup()  {
    println("//Take a guess[1-50]")
    player := &Player{}
    input, playerName := getInput()
    player.name = playerName
    player.guess = input
    _ = generateNumber()
    fmt.Printf("%s guessed: %d\n",player.name, player.guess)
}
