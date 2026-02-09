package main

type Tetromino struct {
	Shape  [4][4]rune
	ID     rune
	Width  int
	Height int
}

type Board struct {
	Grid [][]rune
	Size int
}