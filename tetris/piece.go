package tetris

import (
	"fmt"
	"math/rand"
	"time"
)

type Piece interface {
	InitPiece(color string)
	SmartRotate()
}

func GetTetrisPiece() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(7))
	// create an slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6}
	// concat slice with itself
	numbers = append(numbers, numbers...)
	// shuffle the slice
	shuffle(numbers)

	fmt.Println(numbers)

	// slice of pieces
	pieces := []Piece{}

	// move through the slice

	for _, v := range numbers {
		// switch on the value
		switch v {
		case 0:
			piece := &SPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		case 1:
			piecel := &LPiece{}
			piecel.InitPiece("red")
			pieces = append(pieces, piecel)
		case 2:
			piece := &JPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		case 3:
			piece := &ZPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		case 4:
			piece := &OPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		case 5:
			piece := &LPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		case 6:
			piece := &TPiece{}
			piece.InitPiece("red")
			pieces = append(pieces, piece)
		}
	}

	// print the pieces
	for _, v := range pieces {
		fmt.Println(v)
	}
}

// a function that shuffles a slice
func shuffle(slice []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
