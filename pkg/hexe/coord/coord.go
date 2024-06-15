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

type TCoords[T Coord] interface {
}

type CoordPredicate[T Coord] func(coord T) bool

type Coords[T Coord, TS TCoords[T]] interface {
	~[]T
	Axials() Axials
	Cubes() Cubes
	OddRs() OddRs
	EvenRs() EvenRs
	OddQs() OddQs
	EvenQs() EvenQs
	DoubleWidths() DoubleWidths
	DoubleHeights() DoubleHeights
	Copy() TS
	Sorted() TS
	UnionWith(other TS) TS
	IntersectWith(other TS) TS
	DifferenceWith(other TS) TS
}

type QR[T Coord, TS Coords[T, TS]] interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
	FloodFill(n int, blocked CoordPredicate[T]) TS
}

type QRS[T Coord, TS Coords[T, TS]] interface {
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
	FloodFill(n int, blocked CoordPredicate[T]) TS
}

func castAs[F Coord, T Coord](values []F, convertFunc func(F) T) []T {
	result := make([]T, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = convertFunc(values[i])
	}
	return result
}
