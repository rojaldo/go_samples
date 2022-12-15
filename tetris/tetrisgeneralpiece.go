package tetris

import "fmt"

type TetrisGeneralPiece struct {
	shape     [5][5]bool
	color     string
	pieceType string
	rotation  int
	rotations [4][5][5]bool
}

func (t *TetrisGeneralPiece) InitPiece(color string, pieceType string, shape [5][5]bool) {
	t.color = color
	t.pieceType = pieceType
	t.shape = shape

	switch pieceType {
	case "I":
		t.rotations[0] = shape
		t.rotations[1] = getRotation(shape)
		t.rotations[2] = getRotation(t.rotations[1])
		t.rotations[3] = getRotation(t.rotations[2])
		break
	case "L":
		t.rotations[0] = shape
		t.rotations[1] = getRotation(shape)
		t.rotations[2] = getRotation(t.rotations[1])
		t.rotations[3] = getRotation(t.rotations[2])
		break
	case "J":
		t.rotations[0] = shape
		t.rotations[1] = getRotation(shape)
		t.rotations[2] = getRotation(t.rotations[1])
		t.rotations[3] = getRotation(t.rotations[2])
		break
	default:
		t.rotations[0] = shape
		t.rotations[1] = getRotation(shape)
		t.rotations[2] = getRotation(t.rotations[1])
		t.rotations[3] = getRotation(t.rotations[2])
		break
	}
}

func (t *TetrisGeneralPiece) NewPiece(typePiece int) TetrisGeneralPiece {
	switch typePiece {
	case 0:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, true, false},
				{false, false, false, false, false},
			},
			color:     "blue",
			pieceType: "L",
			rotation:  0}
	case 1:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, true, true, false, false},
				{false, false, false, false, false},
			},
			color:     "orange",
			pieceType: "J",
			rotation:  0}
	case 2:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, true, false, false},
				{false, true, true, true, false},
				{false, false, false, false, false},
			},
			color:     "red",
			pieceType: "T",
			rotation:  0}
	case 3:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, true, true, false},
				{false, false, true, true, false},
				{false, false, false, false, false},
			},
			color:     "yellow",
			pieceType: "O",
			rotation:  0}
	case 4:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, true, true, false},
				{false, true, true, false, false},
				{false, false, false, false, false},
			},
			color:     "green",
			pieceType: "S",
			rotation:  0}
	case 5:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, true, true, false, false},
				{false, false, true, true, false},
				{false, false, false, false, false},
			},
			color:     "purple",
			pieceType: "Z",
			rotation:  0}
	case 6:
		return TetrisGeneralPiece{
			shape: [5][5]bool{
				{false, false, false, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
			},
			color:     "cyan",
			pieceType: "I",
			rotation:  0}
	default:
		return TetrisGeneralPiece{}
	}
}

func (t *TetrisGeneralPiece) SetColor(color string) {
	t.color = color
}

func (t *TetrisGeneralPiece) SetPieceType(pieceType string) {
	t.pieceType = pieceType
}

func (t *TetrisGeneralPiece) SetShape(shape [5][5]bool) {
	t.shape = shape
}

func (t TetrisGeneralPiece) String() string {
	return fmt.Sprintf("Color: %v \nPieceType: %v", t.color, t.pieceType)
}

func (t *TetrisGeneralPiece) Rotate() {
	// Rotate the piece 90 degrees counterclockwise
	newShape := [5][5]bool{}
	for i := range t.shape {
		for j, v := range t.shape[i] {
			newShape[j][5-1-i] = v
		}
	}
	t.shape = newShape
}

func (t *TetrisGeneralPiece) SmartRotate() {
	t.rotation = (t.rotation + 1) % 4
	t.shape = t.rotations[t.rotation]
}

func getRotation(shape [5][5]bool) [5][5]bool {
	newShape := [5][5]bool{}
	for i := range shape {
		for j, v := range shape[i] {
			newShape[j][5-1-i] = v
		}
	}
	return newShape
}

func (t *TetrisGeneralPiece) PrintPiece() {
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
