package hexe

type index [2]int

type Grid[T any] interface {
	Axial() AxialGrid[T]
	Cube() CubeGrid[T]
	EvenQ() EvenQGrid[T]
	EvenR() EvenRGrid[T]
	OddQ() OddQGrid[T]
	OddR() OddRGrid[T]
	DoubleWidth() DoubleWidthGrid[T]
	DoubleHeight() DoubleHeightGrid[T]
}

type grid[T any] struct {
	encoder Encoder[T]
	decoder Decoder[T]
	items   map[index]T
}

// NewGrid returns a new hexagonal grid
func NewGrid[T any](options ...Option[T]) AxialGrid[T] {
	g := &grid[T]{
		items: make(map[index]T),
	}
	for _, opt := range options {
		opt.apply(g)
	}
	return newAxialGrid(g)
}

func (g *grid[T]) get(q int, r int) T {
	return g.items[index{q, r}]
}

func (g *grid[T]) index(q int, r int) (T, bool) {
	value, ok := g.items[index{q, r}]
	return value, ok
}

func (g *grid[T]) set(q int, r int, value T) {
	g.items[index{q, r}] = value
}

func (g *grid[T]) delete(q int, r int) {
	delete(g.items, index{q, r})
}

// Axial exposes grid access methods in axial coordinates
func (g *grid[T]) Axial() AxialGrid[T] {
	return newAxialGrid(g)
}

// Cube exposes grid access methods in cube coordinates
func (g *grid[T]) Cube() CubeGrid[T] {
	return newCubeGrid(g)
}

func (g *grid[T]) EvenQ() EvenQGrid[T] {
	return newEvenQ(g)
}

func (g *grid[T]) EvenR() EvenRGrid[T] {
	return newEvenR(g)
}

func (g *grid[T]) OddQ() OddQGrid[T] {
	return newOddQ(g)
}

func (g *grid[T]) OddR() OddRGrid[T] {
	return newOddR(g)
}

func (g *grid[T]) DoubleWidth() DoubleWidthGrid[T] {
	return newDoubleWidth(g)
}

func (g *grid[T]) DoubleHeight() DoubleHeightGrid[T] {
	return newDoubleHeight(g)
}
