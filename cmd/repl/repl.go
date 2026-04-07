package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("hello world")

	fmt.Println(getWantsToPlay())

	file, err := os.Open(pickTemplate())
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i += 1 {
		fmt.Println(i, scanner.Text())
	}
}

func getWantsToPlay() bool {
	switch sanitizeInput(getInput("Do you want to play the game? (Y/n): ")) {
	case "", "Y", "YES":
		return true
	case "N", "NO":
		return false
	default:
		fmt.Println("Invalid input")
		return getWantsToPlay()
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	// NOTE: scanner.Err?
	return scanner.Text()
}
func sanitizeInput(input string) string {
	return strings.ToUpper(input)
}

// im gonna assume not buggy at all
func pickTemplate() string {
	i := 0

	curPath := ""

	filepath.WalkDir("./res", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		i += 1

		if rand.Intn(i) == 0 {
			curPath = path
		}

		return nil
	})

	return curPath

	// file, _ := os.Open("./res/choose_a_file.txt")
	// return file
}

// regex in: <\s*(.)(.+?)\s*> out:{{ .\u$1$2 }} DAMNNN
// TODO: regex: \{\{\s*(.+?)\s*\}\}
