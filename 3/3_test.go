package main

import (
	"bufio"
	"examples/adventofcode2022/test_utils"
	"strings"
	"testing"
)

func TestRucksacks(t *testing.T) {
	t.Run("split a single rucksacks items into compartments", func(t *testing.T) {
		data := "vJrwpWtwJgWrhcsFMMfFFhFp"
		wanta, wantb := []byte{118, 74, 114, 119, 112, 87, 116, 119, 74, 103, 87, 114},
			[]byte{104, 99, 115, 70, 77, 77, 102, 70, 70, 104, 70, 112}

		gota, gotb := SplitItems(data)

		test_utils.AssertDeepEqual(t, gota, wanta)
		test_utils.AssertDeepEqual(t, gotb, wantb)

	})
	t.Run("get items from a set of rucksack's", func(t *testing.T) {
		data := "vJrwpWtwJgWrhcsFMMfFFhFp\n"
		data += "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n"
		data += "PmmdzqPrVvPwwTWBwg\n"
		data += "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n"
		data += "ttgJtRGJQctTZtZT\n"
		data += "CrZsJsPPZsGzwwsLwLmpwMDw\n"

		want := 157
		want2 := 70
		got, got2, err := ProcessItems(bufio.NewReader(strings.NewReader(data)))

		test_utils.AssertNoError(t, err)
		test_utils.AssertEqual(t, got, want)
		test_utils.AssertEqual(t, got2, want2)

	})
	t.Run("find intersection in the two list of items", func(t *testing.T) {
		dataa, datab := []byte{118, 74, 114, 119, 112, 87, 116, 119, 74, 103, 87, 114},
			[]byte{104, 99, 115, 70, 77, 77, 102, 70, 70, 104, 70, 112}

		want := []byte{112}
		got := FindInterjectItems(dataa, datab)

		test_utils.AssertDeepEqual(t, got, want)
	})
	t.Run("prioritize list of items", func(t *testing.T) {
		data := []byte{112, 76, 80, 118, 116, 115}
		want := []int{16, 38, 42, 22, 20, 19}
		got, err := PrioritizeItems(data)

		test_utils.AssertNoError(t, err)
		test_utils.AssertDeepEqual(t, got, want)
	})
	t.Run("error when not a letter", func(t *testing.T) {
		data := []byte{112, 64, 76, 91, 96, 123}

		_, err := PrioritizeItems(data)

		test_utils.AssertError(t, err)

	})
	t.Run("sum priorities", func(t *testing.T) {
		data := []int{16, 38, 42, 22, 20, 19}
		want := 157
		got := SumItems(data)

		test_utils.AssertEqual(t, got, want)
	})
	t.Run("find interject from 3 arrays", func(t *testing.T) {

	})
}
