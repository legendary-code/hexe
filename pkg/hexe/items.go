package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

// Items represents a map of coordinate to values
type Items[T any, C coord.CCoord] map[C]T
