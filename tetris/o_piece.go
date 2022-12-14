package tetris

type OPiece struct {
	TetrisPiece
	rotation  int
	rotations [1][5][5]bool
}

func (l *OPiece) InitPiece(color string) {
	l.color = color
	l.pieceType = "O"
	l.shape = [5][5]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, true, true, false},
		{false, false, true, true, false},
		{false, false, false, false, false},
	}

	l.rotations[0] = l.shape

}

func (t *OPiece) SmartRotate() {
}
