package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type Items[T any, C coord.Coord, CS coord.Coords[C]] []Item[T, C]

func (i Items[T, C, CS]) Values() []T {
	items := make([]T, len(i))
	for j := 0; j < len(i); j++ {
		items[j] = i[j].Value()
	}
	return items
}

func (i Items[T, C, CS]) Indices() CS {
	items := make(CS, len(i))
	for j := 0; j < len(i); j++ {
		items[j] = i[j].Index()
	}
	return items
}

func castAs[T any, FC coord.Coord, FCS coord.Coords[FC], TC coord.Coord, TCS coord.Coords[TC]](
	items Items[T, FC, FCS],
	castFunc func(c FC) TC,
) Items[T, TC, TCS] {
	castedItems := make(Items[T, TC, TCS], len(items))
	for j := 0; j < len(items); j++ {
		i := items[j]
		castedItems[j] = newItem(castFunc(i.Index()), i.Value())
	}
	return castedItems
}

func (i Items[T, C, CS]) Axials() Items[T, coord.Axial, coord.Axials] {
	return castAs[T, C, CS, coord.Axial, coord.Axials](i, C.Axial)
}

func (i Items[T, C, CS]) Cubes() Items[T, coord.Cube, coord.Cubes] {
	return castAs[T, C, CS, coord.Cube, coord.Cubes](i, C.Cube)
}

func (i Items[T, C, CS]) OddRs() Items[T, coord.OddR, coord.OddRs] {
	return castAs[T, C, CS, coord.OddR, coord.OddRs](i, C.OddR)
}

func (i Items[T, C, CS]) EvenRs() Items[T, coord.EvenR, coord.EvenRs] {
	return castAs[T, C, CS, coord.EvenR, coord.EvenRs](i, C.EvenR)
}

func (i Items[T, C, CS]) OddQs() Items[T, coord.OddQ, coord.OddQs] {
	return castAs[T, C, CS, coord.OddQ, coord.OddQs](i, C.OddQ)
}

func (i Items[T, C, CS]) EvenQs() Items[T, coord.EvenQ, coord.EvenQs] {
	return castAs[T, C, CS, coord.EvenQ, coord.EvenQs](i, C.EvenQ)
}

func (i Items[T, C, CS]) DoubleWidths() Items[T, coord.DoubleWidth, coord.DoubleWidths] {
	return castAs[T, C, CS, coord.DoubleWidth, coord.DoubleWidths](i, C.DoubleWidth)
}

func (i Items[T, C, CS]) DoubleHeights() Items[T, coord.DoubleHeight, coord.DoubleHeights] {
	return castAs[T, C, CS, coord.DoubleHeight, coord.DoubleHeights](i, C.DoubleHeight)
}
