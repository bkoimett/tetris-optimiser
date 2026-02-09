package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	filename := os.Args[1]

	// Parse tetrominoes
	tetrominos, err := parseFile(filename)
	if err != nil {
		fmt.Println("ERROR: couldn't parse files")
		return
	}

	// Solve
	solution := solve(tetrominos)

	// Print solution
	for _, line := range solution {
		fmt.Println(line)
	}
}