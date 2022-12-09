package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

type coord [2]int

func main() {
	lines := strings.Split(inputsContent, "\n")
	forest := [][]int{}
	for _, line := range lines {
		forest = append(forest, strToIntArr(line))
	}

	visible := mapVisibleTrees(forest)
	fmt.Printf("visible trees => %d\n", len(visible))

	scenicScore := findHighestScenicScore(forest)
	fmt.Printf("highest scenic score => %d\n", scenicScore)

}

func mapVisibleTrees(forest [][]int) map[coord]bool {
	visible := make(map[coord]bool)
	rowsLength := len(forest)
	columnsLenght := len(forest[0])

	var tallestTree int

	for row := 0; row < rowsLength; row++ {

		// left to right
		tallestTree = -1
		for col := 0; col < columnsLenght; col++ {
			if forest[row][col] > tallestTree {
				visible[coord{row, col}] = true
				tallestTree = forest[row][col]
			}
		}

		// right to left
		tallestTree = -1
		for col := columnsLenght - 1; col > 0; col-- {
			if forest[row][col] > tallestTree {
				visible[coord{row, col}] = true
				tallestTree = forest[row][col]
			}
		}
	}

	for col := 0; col < columnsLenght; col++ {

		// bottom to top
		tallestTree = -1
		for row := rowsLength - 1; row > 0; row-- {
			if forest[row][col] > tallestTree {
				visible[coord{row, col}] = true
				tallestTree = forest[row][col]
			}
		}

		// top to bottom
		tallestTree = -1
		for row := 0; row < rowsLength; row++ {
			if forest[row][col] > tallestTree {
				visible[coord{row, col}] = true
				tallestTree = forest[row][col]
			}
		}

	}
	return visible
}

func findHighestScenicScore(forest [][]int) int {

	max := 0
	rowsLength := len(forest)
	columnsLenght := len(forest[0])

	for row := 1; row < rowsLength-1; row++ {
		for col := 1; col < columnsLenght-1; col++ {
			// TODO: there must be a  way to DRY this
			left := countLeft(forest, row, col)
			right := countRight(forest, row, col)
			up := countUp(forest, row, col)
			down := countDown(forest, row, col)

			total := left * right * up * down
			if total > max {
				max = total
			}
		}
	}

	return max
}

func countDown(forest [][]int, row, col int) (count int) {
	for slide := row + 1; slide < len(forest); slide++ {
		count++
		if forest[row][col] <= forest[slide][col] {
			break
		}
	}
	return
}

func countUp(forest [][]int, row, col int) (count int) {
	for i := row - 1; i >= 0; i-- {
		count++
		if forest[row][col] <= forest[i][col] {
			break
		}
	}
	return
}

func countRight(forest [][]int, row, col int) (count int) {
	for i := col + 1; i < len(forest); i++ {
		count++
		if forest[row][col] <= forest[row][i] {
			break
		}
	}
	return
}

func countLeft(forest [][]int, row, col int) (count int) {
	for slide := col - 1; slide >= 0; slide-- {
		count++
		if forest[row][col] <= forest[row][slide] {
			break
		}
	}
	return
}

func strToIntArr(str string) (intArr []int) {
	for _, v := range string(str) {
		int, _ := strconv.Atoi(string(v))
		intArr = append(intArr, int)
	}
	return
}
