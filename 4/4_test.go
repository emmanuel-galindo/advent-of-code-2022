package main

import (
	"examples/adventofcode2022/test_utils"
	"testing"
)

func TestIsContaining(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{"an item is not contained", `25-94,24-67`, false},
		{"right contains left", `67-88,41-89`, true},
		{"left contains right", `3-60,4-4`, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			got := IsContaining([]byte(tC.input))

			test_utils.AssertEqual(t, got, tC.want)

		})
	}
}

func TestIsOverlapping(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{"an item is not overlapped", `20-86,95-99`, false},
		{"right overlaps left", `14-83,5-14`, true},
		{"left overlaps right", `6-90,7-99`, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			got := IsOverlapping([]byte(tC.input))

			test_utils.AssertEqual(t, got, tC.want)

		})
	}
}
