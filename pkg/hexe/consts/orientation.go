package consts

// Orientation represents how hexagons are oriented
type Orientation bool

const (
	// FlatTop represents hexagons with flat side facing the top
	FlatTop Orientation = true

	// PointyTop represents hexagons with a pointy corner facing the top
	PointyTop Orientation = false
)
