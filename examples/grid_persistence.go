package main

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"strings"
)

type StringEncoderDecoder struct {
}

func (s *StringEncoderDecoder) Encode(value string) ([]byte, error) {
	return []byte(value), nil
}

func (s *StringEncoderDecoder) Decode(bytes []byte) (string, error) {
	return string(bytes), nil
}

func gridPersistenceExample() {
	codec := &StringEncoderDecoder{}
	grid := hexe.NewAxialGrid[string](
		hexe.WithEncoder[string](codec),
		hexe.WithDecoder[string](codec),
	)

	grid.Set(coord.NewAxial(0, 1), "foo")
	grid.Set(coord.NewAxial(1, 0), "bar")

	sb := strings.Builder{}

	err := grid.Encode(&sb)
	if err != nil {
		panic(err)
	}

	grid.Clear()

	r := strings.NewReader(sb.String())
	err = grid.Decode(r)
	if err != nil {
		panic(err)
	}

	for i := grid.Iterator(); i.Next(); {
		fmt.Printf("%v => %s\n", i.Index(), i.Item())
	}
}
