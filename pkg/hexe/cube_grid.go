package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type CubeGrid[T any] Grid[T, coord.Cube, *coord.Cubes]

type cubeGrid[T any] struct {
	*grid[T, coord.Cube, *coord.Cubes]
}

func NewCubeGrid[T any](options ...Option[T]) CubeGrid[T] {
	return &cubeGrid[T]{
		grid: newGrid[T, coord.Cube, *coord.Cubes](options...),
	}
}

func newCubeGrid[T any, C coord.CCoord, CS coord.CCoords](grid *grid[T, C, CS]) CubeGrid[T] {
	return &cubeGrid[T]{
		grid: convertGrid[T, C, CS, coord.Cube, *coord.Cubes](grid),
	}
}

func (c *cubeGrid[T]) Get(index coord.Cube) T {
	return c.grid.get(index)
}

func (c *cubeGrid[T]) GetAll(indices *coord.Cubes) Items[T, coord.Cube] {
	return c.grid.getAll(indices, coord.Axial.Cube)
}

func (c *cubeGrid[T]) Index(index coord.Cube) (T, bool) {
	return c.grid.index(index)
}

func (c *cubeGrid[T]) Set(index coord.Cube, value T) {
	c.grid.set(index, value)
}

func (c *cubeGrid[T]) Delete(index coord.Cube) {
	c.grid.delete(index)
}

func (c *cubeGrid[T]) Iterator() GridIterator[T, coord.Cube, *coord.Cubes] {
	return c.iterator(coord.Axial.Cube)
}
