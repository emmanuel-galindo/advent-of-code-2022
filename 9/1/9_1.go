package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

// x,y
type coord [2]int

type Head struct {
	coord
}

type Tail struct {
	coord
	visited map[coord]bool
}

var moves = map[rune]coord{
	'L': {-1, 0},
	'R': {+1, 0},
	'U': {0, -1},
	'D': {0, +1},
}

func main() {
	lines := strings.Split(inputsContent, "\n")

	// 0,0 is middle of grid
	head := Head{coord: coord{0, 0}}
	tail := Tail{coord: coord{0, 0}, visited: map[coord]bool{}}

	for _, line := range lines {
		var direction rune
		var qty int
		fmt.Sscanf(line, "%c %d", &direction, &qty)
		for i := 0; i < qty; i++ {
			head.move(direction)
			tail.follow(&head)
		}
	}
	fmt.Printf("%d", len(tail.visited))
}

func (t *Tail) follow(h *Head) {
	xdiff, ydiff := h.coord[0]-t.coord[0], h.coord[1]-t.coord[1]
	if abs(xdiff) > 1 || abs(ydiff) > 1 {
		t.coord = coord{t.coord[0] + sign(xdiff), t.coord[1] + sign(ydiff)}
		t.visited[t.coord] = true
	}
}

func (h *Head) move(direction rune) {
	move := moves[direction]
	h.coord = coord{h.coord[0] + move[0], h.coord[1] + move[1]}
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
