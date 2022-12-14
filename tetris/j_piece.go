package tetris

type JPiece struct {
	TetrisPiece
	rotation  int
	rotations [4][5][5]bool
}

func (l *JPiece) InitPiece(color string) {
	l.color = color
	l.pieceType = "J"
	l.shape = [5][5]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, true, true, false, false},
		{false, false, false, false, false},
	}

	l.rotations[0] = l.shape
	l.rotations[1] = getRotation(l.rotations[0])
	l.rotations[2] = getRotation(l.rotations[1])
	l.rotations[3] = getRotation(l.rotations[2])
}

func (t *JPiece) SmartRotate() {
	t.rotation = (t.rotation + 1) % 4
	t.shape = t.rotations[t.rotation]
}
