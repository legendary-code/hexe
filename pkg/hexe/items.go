package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type Items[T any, C coord.CCoord] map[C]T
