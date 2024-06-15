package coord

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

type Coords[T Coord] interface {
	~[]T
	Axials() Axials
	Cubes() Cubes
	OddRs() OddRs
	EvenRs() EvenRs
	OddQs() OddQs
	EvenQs() EvenQs
	DoubleWidths() DoubleWidths
	DoubleHeights() DoubleHeights
}

type QR[T Coord, TS Coords[T]] interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
}

type QRS[T Coord, TS Coords[T]] interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
}

func castAs[F Coord, T Coord](values []F, convertFunc func(F) T) []T {
	result := make([]T, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = convertFunc(values[i])
	}
	return result
}
