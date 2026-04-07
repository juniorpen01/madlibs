package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	} // NOTE: scanner.Err?
}

// regex in: <\s*(.)(.+?)\s*> out:{{ .\u$1$2 }} DAMNNN
// TODO: regex: \{\{\s*(.+?)\s*\}\}
