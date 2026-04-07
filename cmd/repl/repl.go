package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello world")

	scanner := bufio.NewScanner(os.Stdin)

	switch sanitizeInput(getInput(scanner, "Do you want to play the game? (Y/n): ")) {
	case "", "Y", "YES":
		panic("todo")
	case "N", "NO":
		panic("todo")
	default:
		panic("todo")
	}
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	// NOTE: scanner.Err?
	return scanner.Text()
}
func sanitizeInput(input string) string {
	return strings.ToUpper(input)
}

// regex in: <\s*(.)(.+?)\s*> out:{{ .\u$1$2 }} DAMNNN
// TODO: regex: \{\{\s*(.+?)\s*\}\}
