// Code generated by internal/hexe/gen; DO NOT EDIT.

package coord

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type Axial [2]int

func NewAxial(q int, r int) Axial {
	return Axial{q, r}
}
func ZeroAxial() Axial {
	return Axial{0, 0}
}
func (a Axial) Type() consts.CoordType {
	return consts.Axial
}
func (a Axial) Convert(typ consts.CoordType) Coord {
	return convertCoord(a, typ)
}
func (a Axial) Q() int {
	return a[0]
}
func (a Axial) R() int {
	return a[1]
}
func (a Axial) Unpack() (int, int) {
	return a[0], a[1]
}
func (a Axial) Neighbor(angle int) Axial {
	return a.Cube().Neighbor(angle).Axial()
}
func (a Axial) Add(other Axial) Axial {
	return a.Cube().Add(other.Cube()).Axial()
}
func (a Axial) Scale(factor int) Axial {
	return a.Cube().Scale(factor).Axial()
}
func (a Axial) Neighbors() Axials {
	return a.Cube().Neighbors().Axials()
}
func (a Axial) DiagonalNeighbors() Axials {
	return a.Cube().DiagonalNeighbors().Axials()
}
func (a Axial) DistanceTo(other Axial) int {
	return a.Cube().DistanceTo(other.Cube())
}
func (a Axial) LineTo(other Axial) Axials {
	return a.Cube().LineTo(other.Cube()).Axials()
}
func (a Axial) MovementRange(n int) Axials {
	return a.Cube().MovementRange(n).Axials()
}
func (a Axial) FloodFill(n int, blocked Predicate[Axial]) Axials {
	return a.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.Axial())
	}).Axials()
}
func (a Axial) Rotate(center Axial, angle int) Axial {
	return a.Cube().Rotate(center.Cube(), angle).Axial()
}
func (a Axial) ReflectQ() Axial {
	return a.Cube().ReflectQ().Axial()
}
func (a Axial) ReflectR() Axial {
	return a.Cube().ReflectR().Axial()
}
func (a Axial) ReflectS() Axial {
	return a.Cube().ReflectS().Axial()
}
func (a Axial) Ring(radius int) Axials {
	return a.Cube().Ring(radius).Axials()
}

type DoubleHeight [2]int

func NewDoubleHeight(q int, r int) DoubleHeight {
	return DoubleHeight{q, r}
}
func ZeroDoubleHeight() DoubleHeight {
	return DoubleHeight{0, 0}
}
func (d DoubleHeight) Type() consts.CoordType {
	return consts.DoubleHeight
}
func (d DoubleHeight) Convert(typ consts.CoordType) Coord {
	return convertCoord(d, typ)
}
func (d DoubleHeight) Q() int {
	return d[0]
}
func (d DoubleHeight) R() int {
	return d[1]
}
func (d DoubleHeight) Unpack() (int, int) {
	return d[0], d[1]
}
func (d DoubleHeight) Neighbor(angle int) DoubleHeight {
	return d.Cube().Neighbor(angle).DoubleHeight()
}
func (d DoubleHeight) Add(other DoubleHeight) DoubleHeight {
	return d.Cube().Add(other.Cube()).DoubleHeight()
}
func (d DoubleHeight) Scale(factor int) DoubleHeight {
	return d.Cube().Scale(factor).DoubleHeight()
}
func (d DoubleHeight) Neighbors() DoubleHeights {
	return d.Cube().Neighbors().DoubleHeights()
}
func (d DoubleHeight) DiagonalNeighbors() DoubleHeights {
	return d.Cube().DiagonalNeighbors().DoubleHeights()
}
func (d DoubleHeight) DistanceTo(other DoubleHeight) int {
	return d.Cube().DistanceTo(other.Cube())
}
func (d DoubleHeight) LineTo(other DoubleHeight) DoubleHeights {
	return d.Cube().LineTo(other.Cube()).DoubleHeights()
}
func (d DoubleHeight) MovementRange(n int) DoubleHeights {
	return d.Cube().MovementRange(n).DoubleHeights()
}
func (d DoubleHeight) FloodFill(n int, blocked Predicate[DoubleHeight]) DoubleHeights {
	return d.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.DoubleHeight())
	}).DoubleHeights()
}
func (d DoubleHeight) Rotate(center DoubleHeight, angle int) DoubleHeight {
	return d.Cube().Rotate(center.Cube(), angle).DoubleHeight()
}
func (d DoubleHeight) ReflectQ() DoubleHeight {
	return d.Cube().ReflectQ().DoubleHeight()
}
func (d DoubleHeight) ReflectR() DoubleHeight {
	return d.Cube().ReflectR().DoubleHeight()
}
func (d DoubleHeight) ReflectS() DoubleHeight {
	return d.Cube().ReflectS().DoubleHeight()
}
func (d DoubleHeight) Ring(radius int) DoubleHeights {
	return d.Cube().Ring(radius).DoubleHeights()
}

type DoubleWidth [2]int

func NewDoubleWidth(q int, r int) DoubleWidth {
	return DoubleWidth{q, r}
}
func ZeroDoubleWidth() DoubleWidth {
	return DoubleWidth{0, 0}
}
func (d DoubleWidth) Type() consts.CoordType {
	return consts.DoubleWidth
}
func (d DoubleWidth) Convert(typ consts.CoordType) Coord {
	return convertCoord(d, typ)
}
func (d DoubleWidth) Q() int {
	return d[0]
}
func (d DoubleWidth) R() int {
	return d[1]
}
func (d DoubleWidth) Unpack() (int, int) {
	return d[0], d[1]
}
func (d DoubleWidth) Neighbor(angle int) DoubleWidth {
	return d.Cube().Neighbor(angle).DoubleWidth()
}
func (d DoubleWidth) Add(other DoubleWidth) DoubleWidth {
	return d.Cube().Add(other.Cube()).DoubleWidth()
}
func (d DoubleWidth) Scale(factor int) DoubleWidth {
	return d.Cube().Scale(factor).DoubleWidth()
}
func (d DoubleWidth) Neighbors() DoubleWidths {
	return d.Cube().Neighbors().DoubleWidths()
}
func (d DoubleWidth) DiagonalNeighbors() DoubleWidths {
	return d.Cube().DiagonalNeighbors().DoubleWidths()
}
func (d DoubleWidth) DistanceTo(other DoubleWidth) int {
	return d.Cube().DistanceTo(other.Cube())
}
func (d DoubleWidth) LineTo(other DoubleWidth) DoubleWidths {
	return d.Cube().LineTo(other.Cube()).DoubleWidths()
}
func (d DoubleWidth) MovementRange(n int) DoubleWidths {
	return d.Cube().MovementRange(n).DoubleWidths()
}
func (d DoubleWidth) FloodFill(n int, blocked Predicate[DoubleWidth]) DoubleWidths {
	return d.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.DoubleWidth())
	}).DoubleWidths()
}
func (d DoubleWidth) Rotate(center DoubleWidth, angle int) DoubleWidth {
	return d.Cube().Rotate(center.Cube(), angle).DoubleWidth()
}
func (d DoubleWidth) ReflectQ() DoubleWidth {
	return d.Cube().ReflectQ().DoubleWidth()
}
func (d DoubleWidth) ReflectR() DoubleWidth {
	return d.Cube().ReflectR().DoubleWidth()
}
func (d DoubleWidth) ReflectS() DoubleWidth {
	return d.Cube().ReflectS().DoubleWidth()
}
func (d DoubleWidth) Ring(radius int) DoubleWidths {
	return d.Cube().Ring(radius).DoubleWidths()
}

type EvenQ [2]int

func NewEvenQ(q int, r int) EvenQ {
	return EvenQ{q, r}
}
func ZeroEvenQ() EvenQ {
	return EvenQ{0, 0}
}
func (e EvenQ) Type() consts.CoordType {
	return consts.EvenQ
}
func (e EvenQ) Convert(typ consts.CoordType) Coord {
	return convertCoord(e, typ)
}
func (e EvenQ) Q() int {
	return e[0]
}
func (e EvenQ) R() int {
	return e[1]
}
func (e EvenQ) Unpack() (int, int) {
	return e[0], e[1]
}
func (e EvenQ) Neighbor(angle int) EvenQ {
	return e.Cube().Neighbor(angle).EvenQ()
}
func (e EvenQ) Add(other EvenQ) EvenQ {
	return e.Cube().Add(other.Cube()).EvenQ()
}
func (e EvenQ) Scale(factor int) EvenQ {
	return e.Cube().Scale(factor).EvenQ()
}
func (e EvenQ) Neighbors() EvenQs {
	return e.Cube().Neighbors().EvenQs()
}
func (e EvenQ) DiagonalNeighbors() EvenQs {
	return e.Cube().DiagonalNeighbors().EvenQs()
}
func (e EvenQ) DistanceTo(other EvenQ) int {
	return e.Cube().DistanceTo(other.Cube())
}
func (e EvenQ) LineTo(other EvenQ) EvenQs {
	return e.Cube().LineTo(other.Cube()).EvenQs()
}
func (e EvenQ) MovementRange(n int) EvenQs {
	return e.Cube().MovementRange(n).EvenQs()
}
func (e EvenQ) FloodFill(n int, blocked Predicate[EvenQ]) EvenQs {
	return e.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.EvenQ())
	}).EvenQs()
}
func (e EvenQ) Rotate(center EvenQ, angle int) EvenQ {
	return e.Cube().Rotate(center.Cube(), angle).EvenQ()
}
func (e EvenQ) ReflectQ() EvenQ {
	return e.Cube().ReflectQ().EvenQ()
}
func (e EvenQ) ReflectR() EvenQ {
	return e.Cube().ReflectR().EvenQ()
}
func (e EvenQ) ReflectS() EvenQ {
	return e.Cube().ReflectS().EvenQ()
}
func (e EvenQ) Ring(radius int) EvenQs {
	return e.Cube().Ring(radius).EvenQs()
}

type EvenR [2]int

func NewEvenR(q int, r int) EvenR {
	return EvenR{q, r}
}
func ZeroEvenR() EvenR {
	return EvenR{0, 0}
}
func (e EvenR) Type() consts.CoordType {
	return consts.EvenR
}
func (e EvenR) Convert(typ consts.CoordType) Coord {
	return convertCoord(e, typ)
}
func (e EvenR) Q() int {
	return e[0]
}
func (e EvenR) R() int {
	return e[1]
}
func (e EvenR) Unpack() (int, int) {
	return e[0], e[1]
}
func (e EvenR) Neighbor(angle int) EvenR {
	return e.Cube().Neighbor(angle).EvenR()
}
func (e EvenR) Add(other EvenR) EvenR {
	return e.Cube().Add(other.Cube()).EvenR()
}
func (e EvenR) Scale(factor int) EvenR {
	return e.Cube().Scale(factor).EvenR()
}
func (e EvenR) Neighbors() EvenRs {
	return e.Cube().Neighbors().EvenRs()
}
func (e EvenR) DiagonalNeighbors() EvenRs {
	return e.Cube().DiagonalNeighbors().EvenRs()
}
func (e EvenR) DistanceTo(other EvenR) int {
	return e.Cube().DistanceTo(other.Cube())
}
func (e EvenR) LineTo(other EvenR) EvenRs {
	return e.Cube().LineTo(other.Cube()).EvenRs()
}
func (e EvenR) MovementRange(n int) EvenRs {
	return e.Cube().MovementRange(n).EvenRs()
}
func (e EvenR) FloodFill(n int, blocked Predicate[EvenR]) EvenRs {
	return e.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.EvenR())
	}).EvenRs()
}
func (e EvenR) Rotate(center EvenR, angle int) EvenR {
	return e.Cube().Rotate(center.Cube(), angle).EvenR()
}
func (e EvenR) ReflectQ() EvenR {
	return e.Cube().ReflectQ().EvenR()
}
func (e EvenR) ReflectR() EvenR {
	return e.Cube().ReflectR().EvenR()
}
func (e EvenR) ReflectS() EvenR {
	return e.Cube().ReflectS().EvenR()
}
func (e EvenR) Ring(radius int) EvenRs {
	return e.Cube().Ring(radius).EvenRs()
}

type OddQ [2]int

func NewOddQ(q int, r int) OddQ {
	return OddQ{q, r}
}
func ZeroOddQ() OddQ {
	return OddQ{0, 0}
}
func (o OddQ) Type() consts.CoordType {
	return consts.OddQ
}
func (o OddQ) Convert(typ consts.CoordType) Coord {
	return convertCoord(o, typ)
}
func (o OddQ) Q() int {
	return o[0]
}
func (o OddQ) R() int {
	return o[1]
}
func (o OddQ) Unpack() (int, int) {
	return o[0], o[1]
}
func (o OddQ) Neighbor(angle int) OddQ {
	return o.Cube().Neighbor(angle).OddQ()
}
func (o OddQ) Add(other OddQ) OddQ {
	return o.Cube().Add(other.Cube()).OddQ()
}
func (o OddQ) Scale(factor int) OddQ {
	return o.Cube().Scale(factor).OddQ()
}
func (o OddQ) Neighbors() OddQs {
	return o.Cube().Neighbors().OddQs()
}
func (o OddQ) DiagonalNeighbors() OddQs {
	return o.Cube().DiagonalNeighbors().OddQs()
}
func (o OddQ) DistanceTo(other OddQ) int {
	return o.Cube().DistanceTo(other.Cube())
}
func (o OddQ) LineTo(other OddQ) OddQs {
	return o.Cube().LineTo(other.Cube()).OddQs()
}
func (o OddQ) MovementRange(n int) OddQs {
	return o.Cube().MovementRange(n).OddQs()
}
func (o OddQ) FloodFill(n int, blocked Predicate[OddQ]) OddQs {
	return o.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.OddQ())
	}).OddQs()
}
func (o OddQ) Rotate(center OddQ, angle int) OddQ {
	return o.Cube().Rotate(center.Cube(), angle).OddQ()
}
func (o OddQ) ReflectQ() OddQ {
	return o.Cube().ReflectQ().OddQ()
}
func (o OddQ) ReflectR() OddQ {
	return o.Cube().ReflectR().OddQ()
}
func (o OddQ) ReflectS() OddQ {
	return o.Cube().ReflectS().OddQ()
}
func (o OddQ) Ring(radius int) OddQs {
	return o.Cube().Ring(radius).OddQs()
}

type OddR [2]int

func NewOddR(q int, r int) OddR {
	return OddR{q, r}
}
func ZeroOddR() OddR {
	return OddR{0, 0}
}
func (o OddR) Type() consts.CoordType {
	return consts.OddR
}
func (o OddR) Convert(typ consts.CoordType) Coord {
	return convertCoord(o, typ)
}
func (o OddR) Q() int {
	return o[0]
}
func (o OddR) R() int {
	return o[1]
}
func (o OddR) Unpack() (int, int) {
	return o[0], o[1]
}
func (o OddR) Neighbor(angle int) OddR {
	return o.Cube().Neighbor(angle).OddR()
}
func (o OddR) Add(other OddR) OddR {
	return o.Cube().Add(other.Cube()).OddR()
}
func (o OddR) Scale(factor int) OddR {
	return o.Cube().Scale(factor).OddR()
}
func (o OddR) Neighbors() OddRs {
	return o.Cube().Neighbors().OddRs()
}
func (o OddR) DiagonalNeighbors() OddRs {
	return o.Cube().DiagonalNeighbors().OddRs()
}
func (o OddR) DistanceTo(other OddR) int {
	return o.Cube().DistanceTo(other.Cube())
}
func (o OddR) LineTo(other OddR) OddRs {
	return o.Cube().LineTo(other.Cube()).OddRs()
}
func (o OddR) MovementRange(n int) OddRs {
	return o.Cube().MovementRange(n).OddRs()
}
func (o OddR) FloodFill(n int, blocked Predicate[OddR]) OddRs {
	return o.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.OddR())
	}).OddRs()
}
func (o OddR) Rotate(center OddR, angle int) OddR {
	return o.Cube().Rotate(center.Cube(), angle).OddR()
}
func (o OddR) ReflectQ() OddR {
	return o.Cube().ReflectQ().OddR()
}
func (o OddR) ReflectR() OddR {
	return o.Cube().ReflectR().OddR()
}
func (o OddR) ReflectS() OddR {
	return o.Cube().ReflectS().OddR()
}
func (o OddR) Ring(radius int) OddRs {
	return o.Cube().Ring(radius).OddRs()
}