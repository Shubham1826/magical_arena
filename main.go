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

	// Read player name
	fmt.Printf("Enter the name of Player %s: ", pos)
	name := strings.TrimSpace(readLine(reader))

	// Read player attributes
	fmt.Printf("Enter %s's attributes:\n", name)
	health := readPositiveInteger(reader, "Health")
	strength := readPositiveInteger(reader, "Strength")
	attack := readPositiveInteger(reader, "Attack")

	return NewPlayer(name, health, strength, attack)
}

// readPositiveInteger reads and validates a positive integer input from the user.
func readPositiveInteger(reader *bufio.Reader, attribute string) int {
	for {
		fmt.Printf("%s: ", attribute)
		input := strings.TrimSpace(readLine(reader))
		value, err := strconv.Atoi(input)
		if err != nil || value <= 0 {
			fmt.Println("❌ Invalid input! Please enter a positive integer.")
			continue
		}
		return value
	}
}

// readLine reads a single line of input from the user.
func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Println("🏟️ Welcome to the Magical Arena Game!")

	// Create Player A
	playerA := readPlayerInput("A")

	// Create Player B
	playerB := readPlayerInput("B")

	// Start the match
	match := NewMatch(playerA, playerB)
	match.Start()
}
