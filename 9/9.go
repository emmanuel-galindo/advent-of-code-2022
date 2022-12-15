package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

var moves = map[rune]coord{
	'L': {-1, 0},
	'R': {+1, 0},
	'U': {0, -1},
	'D': {0, +1},
}

// x,y
type coord [2]int

type Knot struct {
	coord
	visited map[coord]bool
}

func (k *Knot) Follow(prev *Knot) {
	xdiff, ydiff := prev.coord[0]-k.coord[0], prev.coord[1]-k.coord[1]
	if abs(xdiff) > 1 || abs(ydiff) > 1 {
		k.coord = coord{k.coord[0] + sign(xdiff), k.coord[1] + sign(ydiff)}
		k.visited[k.coord] = true
	}
}

func (k *Knot) Move(direction rune) {
	move := moves[direction]
	k.coord = coord{k.coord[0] + move[0], k.coord[1] + move[1]}
}

func main() {
	lines := strings.Split(inputsContent, "\n")

	knots1 := Simulate(lines, 2)
	fmt.Printf("%d\n", len(knots1[1].visited))

	knots2 := Simulate(lines, 10)
	fmt.Printf("%d\n", len(knots2[9].visited))
}

func Simulate(lines []string, knotsQty int) []*Knot {
	// init 10 knots
	var knots []*Knot
	initCoord := coord{0, 0}
	for i := 0; i < knotsQty; i++ {
		knots = append(knots, &Knot{coord: initCoord, visited: map[coord]bool{initCoord: true}})
	}

	for _, line := range lines {
		var direction rune
		var qty int
		fmt.Sscanf(line, "%c %d", &direction, &qty)
		head := knots[0]
		for i := 0; i < qty; i++ {
			head.Move(direction)
			for j := 1; j < knotsQty; j++ {
				knots[j].Follow(knots[j-1])
			}
		}
	}
	return knots
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
