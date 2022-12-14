package tetris

type IPiece struct {
	TetrisPiece
	rotation  int
	rotations [2][5][5]bool
}

func (l *IPiece) InitPiece(color string) {
	l.color = color
	l.pieceType = "I"
	l.shape = [5][5]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, true, true, false},
		{false, true, true, false, false},
		{false, false, false, false, false},
	}

	l.rotations[0] = l.shape
	l.rotations[1] = getRotation(l.rotations[0])
}

func (t *IPiece) SmartRotate() {
	t.rotation = (t.rotation + 1) % 2
	t.shape = t.rotations[t.rotation]
}
