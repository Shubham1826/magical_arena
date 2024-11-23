package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readPlayerInput prompts the user to input player attributes.
func readPlayerInput(pos string) *Player {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter name of playe %s: ", pos)
	name := strings.TrimSpace(readLine(reader))

	fmt.Printf("Enter %s's attributes\n", name)

	fmt.Print("Health: ")
	health, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	fmt.Print("Strength: ")
	strength, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	fmt.Print("Attack: ")
	attack, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	return NewPlayer(name, health, strength, attack)
}

// readLine reads a single line of input from the user.
func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Println("Welcome to the Magical Arena Game!")

	// Create Player A
	playerA := readPlayerInput("A")

	// Create Player B
	playerB := readPlayerInput("B")

	// Start the match
	match := NewMatch(playerA, playerB)
	match.Start()
}
