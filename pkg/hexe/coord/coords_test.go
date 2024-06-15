package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoords_Coords(t *testing.T) {
	t.Parallel()

	for _, typeA := range consts.CoordinateTypes() {
		for _, typeB := range consts.CoordinateTypes() {
			if typeA == typeB {
				continue
			}

			t.Run(fmt.Sprintf("Test%s_%s", typeA.Name(), typeB.Name()), func(t2 *testing.T) {
				t2.Parallel()
				typeACoords := testCoordinates[typeA]
				typeBCoords := testCoordinates[typeB]

				for i, typeACoord := range typeACoords {
					typeBCoord := typeBCoords[i]
					assert.Equal(t2, typeBCoord, typeACoord.Convert(typeB))
				}
			})
		}
	}
}
