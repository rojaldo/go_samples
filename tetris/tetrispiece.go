package tetris

import "fmt"

type TetrisPiece struct {
	shape     [5][5]bool
	color     string
	pieceType string
}

func (t *TetrisPiece) InitPiece(color string, pieceType string, shape [5][5]bool) {
	t.color = color
	t.pieceType = pieceType
	t.shape = shape

}

func (t *TetrisPiece) SetColor(color string) {
	t.color = color
}

func (t *TetrisPiece) SetPieceType(pieceType string) {
	t.pieceType = pieceType
}

func (t *TetrisPiece) SetShape(shape [5][5]bool) {
	t.shape = shape
}

func (t TetrisPiece) String() string {
	return fmt.Sprintf("Color: %v \nPieceType: %v \nShape: %v", t.color, t.pieceType, t.shape)
}

func (t *TetrisPiece) Rotate() {
	// Rotate the piece 90 degrees counterclockwise
	newShape := [5][5]bool{}
	for i := range t.shape {
		for j, v := range t.shape[i] {
			newShape[j][5-1-i] = v
		}
	}
	t.shape = newShape
}

func (t *TetrisPiece) PrintPiece() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if t.shape[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
