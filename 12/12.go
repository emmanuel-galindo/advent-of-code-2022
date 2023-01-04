package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

type Coord struct {
	X        int
	Y        int
	Prev     *Coord
	Visited  bool
	Distance int
	Mark     rune
	Start    bool
	End      bool
}

func run() error {

	// Part 1
	heightmap := ParseMap()
	source1 := Find(heightmap, func(c Coord) bool { return c.Start })
	c := func(to, from Coord) bool { return check(heightmap, to, from) }
	BFS(source1, heightmap, c)
	target1 := Find(heightmap, func(c Coord) bool { return c.End })

	fmt.Printf("Part 1: Shortest path from S to E has %d steps\n", target1.Distance)

	// Part 2
	heightmap = ParseMap()
	source2 := Find(heightmap, func(c Coord) bool { return c.End })
	c = func(to, from Coord) bool { return check(heightmap, from, to) }
	BFS(source2, heightmap, c)
	target2 := FindSmallest(heightmap, 'a')

	fmt.Printf("Part 2: Shortest path from E to nearest a has %d steps", target2.Distance)

	// uncomment to display the grid with the shortest path
	// track(heightmap, target, source)

	return nil
}

func BFS(source Coord, heightmap [][]Coord, c func(Coord, Coord) bool) [][]Coord {
	source.Distance = 0
	queue := []Coord{source}
	var curr Coord

	for len(queue) > 0 {

		curr, queue = queue[0], queue[1:]

		adjs := GetAdjacents(heightmap, curr, c)
		for _, adj := range adjs {
			if adj.Distance > curr.Distance+1 {
				adj.Distance = curr.Distance + 1
				adj.Prev = &heightmap[curr.Y][curr.X]

				queue = append(queue, adj)

				heightmap[adj.Y][adj.X] = adj
			}
		}
	}
	return heightmap
}

// ParseMap prepares the grid. It identifies the start/end marks and replaces the character to their assigned elevation.
// S has an elevation of a, and E and elevation of Z
func ParseMap() [][]Coord {
	lines := strings.Split(inputsContent, "\n")
	heightmap := [][]Coord{}

	for y, line := range lines {
		lineArr := []Coord{}
		for x, mark := range line {
			start, end := false, false
			if mark == 'S' {
				mark = 'a'
				start = true
			}
			if mark == 'E' {
				mark = 'z'
				end = true
			}
			lineArr = append(lineArr, Coord{X: x, Y: y, Distance: math.MaxInt32, Mark: mark, Start: start, End: end})
		}
		heightmap = append(heightmap, lineArr)
	}
	return heightmap
}

func Find(heightmap [][]Coord, fn func(Coord) bool) Coord {
	for _, line := range heightmap {
		for _, coord := range line {
			if fn(coord) {
				return coord
			}
		}
	}
	return Coord{}
}

func FindSmallest(heightmap [][]Coord, target rune) Coord {
	var targetcoord Coord = Coord{Distance: math.MaxInt32}
	for _, line := range heightmap {
		for _, coord := range line {
			if coord.Mark == target && coord.Distance < targetcoord.Distance {
				targetcoord = coord
			}
		}
	}
	return targetcoord
}

var directions = map[rune]Coord{
	'L': {Y: 0, X: -1},
	'R': {Y: 0, X: 1},
	'U': {Y: -1, X: 0},
	'D': {Y: 1, X: 0},
}

func GetAdjacents(heightmap [][]Coord, curr Coord, c func(Coord, Coord) bool) []Coord {
	adjs := []Coord{}

	for direction := range directions {
		targetcoord, err := get(heightmap, curr, direction)
		if err != nil {
			continue
		}

		if c(curr, targetcoord) {
			adjs = append(adjs, targetcoord)
		}
	}
	return adjs
}

func check(heightmap [][]Coord, curr, targetcoord Coord) bool {
	targety := targetcoord.Y
	targetx := targetcoord.X

	curry := curr.Y
	currx := curr.X

	return !(int(heightmap[targety][targetx].Mark) > int(heightmap[curry][currx].Mark)+1)
}

func get(heightmap [][]Coord, curr Coord, orientation rune) (Coord, error) {
	direction := directions[orientation]
	targety := curr.Y + direction.Y
	targetx := curr.X + direction.X
	if targety < 0 || targety > len(heightmap)-1 || targetx < 0 || targetx > len(heightmap[0])-1 {
		return Coord{}, errors.New("direction not valid")
	}
	return heightmap[targety][targetx], nil
}

func showPath(heightmap [][]Coord) {
	for _, line := range heightmap {
		for _, coord := range line {
			if !coord.Visited {
				fmt.Printf("%c", coord.Mark)
			} else {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}
}

func track(heightmap [][]Coord, source, target Coord) {
	q := []Coord{heightmap[source.Y][source.X]}
	var curr Coord
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		heightmap[curr.Y][curr.X].Visited = true
		if curr.Prev.Mark == target.Mark {
			showPath(heightmap)
		}
		q = append(q, heightmap[curr.Prev.Y][curr.Prev.X])
	}
}
