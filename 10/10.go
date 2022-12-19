package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

func main() {
	lines := strings.Split(inputsContent, "\n")

	cycles, x, sumx := 0, 1, 0
	// pixel := [][40]string{}
	pixelrow := []rune{}
	for _, line := range lines {
		var qty int
		runOneCycle(&cycles, &pixelrow, x, &sumx)
		if n, _ := fmt.Sscanf(line, "addx %d", &qty); n > 0 {

			runOneCycle(&cycles, &pixelrow, x, &sumx)

			x += qty
		}
	}
	fmt.Printf("%d\n", sumx)
}

func runOneCycle(cycles *int, pixelrow *[]rune, x int, sumx *int) {
	*cycles++
	printPixel(x, pixelrow)

	CheckAndSumTargetSignalStrength(*cycles, sumx, x)
	CheckAndPrintRow(*cycles, pixelrow)
}

func CheckAndPrintRow(cycles int, pixelrow *[]rune) {
	if cycles%40 == 0 {
		fmt.Printf("%s\n", string(*pixelrow))
		*pixelrow = []rune{}
	}
}

func CheckAndSumTargetSignalStrength(cycles int, sumx *int, x int) {
	if isTargetSignalStrength(cycles) {
		*sumx += x * (cycles)
	}
}

func printPixel(x int, pixelrow *[]rune) {
	if x-1 <= len(*pixelrow) && len(*pixelrow) <= x+1 {
		*pixelrow = append(*pixelrow, '#')
	} else {
		*pixelrow = append(*pixelrow, '.')
	}
}

func isTargetSignalStrength(cycle int) bool {
	return (cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220)
}
