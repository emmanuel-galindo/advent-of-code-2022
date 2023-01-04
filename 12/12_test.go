package main_test

import (
	_ "embed"
	main "examples/adventofcode2022/12"
	"examples/adventofcode2022/test_utils"
	"testing"
)

//go:embed inputs_test.txt
var inputsContent string

func TestAdjacents(t *testing.T) {
	hm := main.ParseMap()
	// TODO: table test
	t.Run("corner left", func(t *testing.T) {

		source := hm[0][0]
		// hm[source.Y][source.X] = '`'

		got := main.GetAdjacents(hm, source)
		want := []main.Coord{{Y: 1, X: 0, Distance: -1}, {Y: 0, X: 1, Distance: -1}}

		test_utils.AssertEqual(t, got[0].Y, want[0].Y)
		test_utils.AssertEqual(t, got[0].X, want[0].X)
		test_utils.AssertEqual(t, got[1].Y, want[1].Y)
		test_utils.AssertEqual(t, got[1].X, want[1].X)
	})
	t.Run("corner right", func(t *testing.T) {

		source := hm[0][7]
		// hm[source.Y][source.X] = '`'

		got := main.GetAdjacents(hm, source)
		want := []main.Coord{{Y: 1, X: 7, Distance: -1}, {Y: 0, X: 6, Distance: -1}}

		test_utils.AssertEqual(t, got[0].Y, want[0].Y)
		test_utils.AssertEqual(t, got[0].X, want[0].X)
		test_utils.AssertEqual(t, got[1].Y, want[1].Y)
		test_utils.AssertEqual(t, got[1].X, want[1].X)
	})

}
