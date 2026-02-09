package main

import (
	"math"
	// "strings"
)

func solve(tetrominos []Tetromino) []string {
	// Calculate minimum board size
	minSize := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))

	for size := minSize; size <= minSize*2; size++ {
		board := makeBoard(size)
		if backtrack(board, tetrominos, 0) {
			return boardToLines(board)
		}
	}

	// If not found, return empty
	return []string{}
}

func makeBoard(size int) [][]rune {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

func backtrack(board [][]rune, tetrominos []Tetromino, index int) bool {
	if index >= len(tetrominos) {
		return true
	}

	tet := tetrominos[index]
	size := len(board)

	for y := 0; y <= size-tet.Height; y++ {
		for x := 0; x <= size-tet.Width; x++ {
			if canPlace(board, tet, x, y) {
				place(board, tet, x, y)
				if backtrack(board, tetrominos, index+1) {
					return true
				}
				remove(board, tet, x, y)
			}
		}
	}
	return false
}

func canPlace(board [][]rune, tet Tetromino, x, y int) bool {
	for ty := 0; ty < tet.Height; ty++ {
		for tx := 0; tx < tet.Width; tx++ {
			if tet.Shape[ty][tx] == '#' {
				if board[y+ty][x+tx] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func place(board [][]rune, tet Tetromino, x, y int) {
	for ty := 0; ty < tet.Height; ty++ {
		for tx := 0; tx < tet.Width; tx++ {
			if tet.Shape[ty][tx] == '#' {
				board[y+ty][x+tx] = tet.ID
			}
		}
	}
}

func remove(board [][]rune, tet Tetromino, x, y int) {
	for ty := 0; ty < tet.Height; ty++ {
		for tx := 0; tx < tet.Width; tx++ {
			if tet.Shape[ty][tx] == '#' {
				board[y+ty][x+tx] = '.'
			}
		}
	}
}

func boardToLines(board [][]rune) []string {
	lines := make([]string, len(board))
	for i, row := range board {
		lines[i] = string(row)
	}
	return lines
}