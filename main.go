package main

import (
	"sample/tetris"
)

func main() {

	uShape := [5][5]bool{
		{false, false, false, false, false},
		{false, false, true, true, false},
		{false, false, true, false, false},
		{false, false, true, true, false},
		{false, false, false, false, false},
	}

	var piece tetris.TetrisGeneralPiece
	piece.InitPiece("red", "U", uShape)
	piece.PrintPiece()
	piece.Rotate()
	piece.PrintPiece()
	var piece2 tetris.LPiece
	piece2.InitPiece("blue")
	piece2.PrintPiece()
	piece2.SmartRotate()
	piece2.PrintPiece()
	var piece3 tetris.OPiece
	piece3.InitPiece("green")
	piece3.PrintPiece()
	piece3.SmartRotate()
	piece3.PrintPiece()

}
