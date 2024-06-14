package coord

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type Coord interface {
	Axial() Axial
	Cube() Cube
	OddR() OddR
	EvenR() EvenR
	OddQ() OddQ
	EvenQ() EvenQ
	DoubleWidth() DoubleWidth
	DoubleHeight() DoubleHeight
}

type QR[T Coord] interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
	Neighbors() [consts.Sides]T
	DiagonalNeighbors() [consts.Sides]T
	DistanceTo(other T) int
}

type QRS[T Coord] interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
	Neighbors() [consts.Sides]T
	DiagonalNeighbors() [consts.Sides]T
	DistanceTo(other T) int
}

type Axial [2]int
type Cube [3]int
type DoubleWidth [2]int
type DoubleHeight [2]int
type EvenQ [2]int
type EvenR [2]int
type OddQ [2]int
type OddR [2]int
