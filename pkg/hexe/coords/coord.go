package coords

type Coord interface {
	Axial() CoordQR
	Cube() CoordQRS
	OddR() CoordQR
	EvenR() CoordQR
	OddQ() CoordQR
	EvenQ() CoordQR
	DoubleWidth() CoordQR
	DoubleHeight() CoordQR
}

type CoordQR interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
}

type CoordQRS interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
}
