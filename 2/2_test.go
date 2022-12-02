package main

import (
	"bufio"
	"examples/adventofcode2022/test_utils"
	"strings"
	"testing"
)

func TestRSCGame(t *testing.T) {
	t.Run("get rounds", func(t *testing.T) {
		data := "A Y\nB X\nC Z"
		want := [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}

		got, _ := GetRounds(bufio.NewReader(strings.NewReader(data)))

		test_utils.AssertDeepEqual(t, got, want)

	})
	t.Run("calculate first round", func(t *testing.T) {
		data := []string{"A", "Y"}
		want := 8
		want2 := 4

		got, got2 := ProcessRound(data)

		test_utils.AssertEqual(t, got, want)
		test_utils.AssertEqual(t, got2, want2)
	})
	t.Run("calculate a set of 3 round", func(t *testing.T) {
		data := [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}
		want := 15
		want2 := 12

		got, got2 := ProcessRounds(data)

		test_utils.AssertEqual(t, got, want)
		test_utils.AssertEqual(t, got2, want2)

	})

}
