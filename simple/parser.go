package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) ([]Tetromino, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetrominos []Tetromino
	var currentPiece []string
	pieceID := 'A'

	for scanner.Scan() {
		line := scanner.Text()
		
		if len(line) == 0 {
			if len(currentPiece) == 4 {
				tet, err := createTetromino(currentPiece, pieceID)
				if err != nil {
					return nil, err
				}
				tetrominos = append(tetrominos, tet)
				pieceID++
				currentPiece = nil
			}
			continue
		}

		if len(currentPiece) < 4 {
			currentPiece = append(currentPiece, line)
		}

		if len(currentPiece) == 4 {
			tet, err := createTetromino(currentPiece, pieceID)
			if err != nil {
				return nil, err
			}
			tetrominos = append(tetrominos, tet)
			pieceID++
			currentPiece = nil
		}
	}

	// Handle last piece
	if len(currentPiece) == 4 {
		tet, err := createTetromino(currentPiece, pieceID)
		if err != nil {
			return nil, err
		}
		tetrominos = append(tetrominos, tet)
	}

	if len(tetrominos) == 0 {
		return nil, fmt.Errorf("ERROR")
	}

	return tetrominos, nil
}

func createTetromino(lines []string, id rune) (Tetromino, error) {
	// Validate exactly 4 lines
	if len(lines) != 4 {
		return Tetromino{}, fmt.Errorf("ERROR")
	}

	// Count # and validate characters
	hashCount := 0
	var shape [4][4]rune

	for i, line := range lines {
		if len(line) != 4 {
			return Tetromino{}, fmt.Errorf("ERROR")
		}
		for j, ch := range line {
			if ch != '#' && ch != '.' {
				return Tetromino{}, fmt.Errorf("ERROR")
			}
			shape[i][j] = ch
			if ch == '#' {
				hashCount++
			}
		}
	}

	// Must have exactly 4 blocks
	if hashCount != 4 {
		return Tetromino{}, fmt.Errorf("ERROR")
	}

	// Check connectivity
	if !isValidTetromino(shape) {
		return Tetromino{}, fmt.Errorf("ERROR")
	}

	// Find bounding box
	minX, minY, maxX, maxY := 3, 3, 0, 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if shape[y][x] == '#' {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	// Create trimmed shape
	var trimmedShape [4][4]rune
	for i := range trimmedShape {
		for j := range trimmedShape[i] {
			trimmedShape[i][j] = '.'
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if shape[y][x] == '#' {
				trimmedShape[y-minY][x-minX] = '#'
			}
		}
	}

	return Tetromino{
		Shape:  trimmedShape,
		ID:     id,
		Width:  width,
		Height: height,
	}, nil
}

func isValidTetromino(shape [4][4]rune) bool {
	// Find first #
	startX, startY := -1, -1
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if shape[y][x] == '#' {
				startX, startY = x, y
				break
			}
		}
		if startX != -1 {
			break
		}
	}

	if startX == -1 {
		return false
	}

	// DFS to count connected blocks
	visited := [4][4]bool{}
	stack := [][2]int{{startX, startY}}
	count := 0

	for len(stack) > 0 {
		x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		if x < 0 || x >= 4 || y < 0 || y >= 4 || visited[y][x] || shape[y][x] != '#' {
			continue
		}

		visited[y][x] = true
		count++

		// Add neighbors
		stack = append(stack, [2]int{x + 1, y})
		stack = append(stack, [2]int{x - 1, y})
		stack = append(stack, [2]int{x, y + 1})
		stack = append(stack, [2]int{x, y - 1})
	}

	return count == 4
}