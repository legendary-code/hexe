package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

//go:generate go run ../../internal/hexe/gen/coords
//go:generate go fmt .

type Grid[T any, C coord.CCoord, CS coord.CCoords] interface {
	Axial() AxialGrid[T]
	Cube() CubeGrid[T]
	EvenQ() EvenQGrid[T]
	EvenR() EvenRGrid[T]
	OddQ() OddQGrid[T]
	OddR() OddRGrid[T]
	DoubleWidth() DoubleWidthGrid[T]
	DoubleHeight() DoubleHeightGrid[T]

	// Get returns item at given coordinate
	Get(index C) T

	// GetAll returns items for given set of coordinates
	GetAll(indices CS) Items[T, C]

	// Index returns item at given coordinate and a bool of whether it exists, similar to indexing into a map
	Index(index C) (T, bool)

	// Set sets an item at the given coordinate
	Set(index C, value T)

	// Delete deletes an item at the given coordinate
	Delete(index C)
}

type grid[T any, C coord.CCoord, CS coord.CCoords] struct {
	*configurableGrid[T]
}

type configurableGrid[T any] struct {
	encoder Encoder[T]
	decoder Decoder[T]
	items   map[coord.Axial]T
}

func newGrid[T any, C coord.CCoord, CS coord.CCoords](options ...Option[T]) *grid[T, C, CS] {
	g := &grid[T, C, CS]{
		configurableGrid: &configurableGrid[T]{
			items: make(map[coord.Axial]T),
		},
	}
	for _, opt := range options {
		opt.apply(g.configurableGrid)
	}
	return g
}

func convertGrid[T any, FC coord.CCoord, FCS coord.CCoords, TC coord.CCoord, TCS coord.CCoords](
	g *grid[T, FC, FCS],
) *grid[T, TC, TCS] {
	return &grid[T, TC, TCS]{
		configurableGrid: g.configurableGrid,
	}
}

func (g *grid[T, C, CS]) get(index C) T {
	return g.items[index.Axial()]
}

func (g *grid[T, C, CS]) getAll(indices CS) Items[T, C] {
	return nil
}

func (g *grid[T, C, CS]) index(index C) (T, bool) {
	value, ok := g.items[index.Axial()]
	return value, ok
}

func (g *grid[T, C, CS]) set(index C, value T) {
	g.items[index.Axial()] = value
}

func (g *grid[T, C, CS]) delete(index C) {
	delete(g.items, index.Axial())
}

// Axial exposes grid access methods in axial coordinates
func (g *grid[T, C, CS]) Axial() AxialGrid[T] {
	return newAxialGrid(g)
}

// Cube exposes grid access methods in cube coordinates
func (g *grid[T, C, CS]) Cube() CubeGrid[T] {
	return newCubeGrid(g)
}

func (g *grid[T, C, CS]) DoubleHeight() DoubleHeightGrid[T] {
	return newDoubleHeightGrid(g)
}

func (g *grid[T, C, CS]) DoubleWidth() DoubleWidthGrid[T] {
	return newDoubleWidthGrid(g)
}

func (g *grid[T, C, CS]) EvenQ() EvenQGrid[T] {
	return newEvenQGrid(g)
}

func (g *grid[T, C, CS]) EvenR() EvenRGrid[T] {
	return newEvenRGrid(g)
}

func (g *grid[T, C, CS]) OddQ() OddQGrid[T] {
	return newOddQGrid(g)
}

func (g *grid[T, C, CS]) OddR() OddRGrid[T] {
	return newOddRGrid(g)
}
