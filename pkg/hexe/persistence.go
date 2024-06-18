package hexe

import (
	"encoding/binary"
	"errors"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"io"
)

type item struct {
	Q   int64
	R   int64
	Len int32
}

func (g *grid[T, C, CS]) Encode(w io.Writer) error {
	if g.encoder == nil {
		return errors.New("grid does not have an encoder configured")
	}

	for c, v := range g.items {
		encoded, err := g.encoder.Encode(v)
		if err != nil {
			return err
		}

		i := &item{Q: int64(c[0]), R: int64(c[1]), Len: int32(len(encoded))}
		err = binary.Write(w, byteOrder, i)
		if err != nil {
			return err
		}
		_, err = w.Write(encoded)
	}

	return nil
}

func (g *grid[T, C, CS]) Decode(r io.Reader) error {
	if g.decoder == nil {
		return errors.New("grid does not have a decoder configured")
	}

	for {
		var i item
		err := binary.Read(r, byteOrder, &i)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		data := make([]byte, i.Len)
		_, err = r.Read(data)
		if err != nil {
			return err
		}

		value, err := g.decoder.Decode(data)
		if err != nil {
			return err
		}

		g.items[coord.NewAxial(int(i.Q), int(i.R))] = value
	}

	return nil
}
