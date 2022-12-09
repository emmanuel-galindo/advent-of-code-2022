package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

const (
	totalSpace     = 70_000_000
	updateReqSpace = 30_000_000
	maxDirSize     = 100_000
)

type Dir struct {
	dirs   map[string]*Dir
	files  *[]int
	prev   *Dir
	total  int
	rtotal int
}

func main() {
	lines := strings.Split(inputsContent, "\n")

	// Create the tree
	root := &Dir{prev: nil}
	processCommand(lines, 1, root)

	// Part 1
	var total int = 0
	recursiveSum(root, &total)
	fmt.Printf("sum of dirs with recursive total less than < %d => %d\n", maxDirSize, total)

	// Part 2
	target := updateReqSpace - (totalSpace - root.rtotal)
	smallerCandidate := root.rtotal
	collectHigherThan(root, target, &smallerCandidate)
	fmt.Printf("The size of the smallest dir to free enough space (%d) for the update is %d", target, smallerCandidate)
}

func collectHigherThan(currDir *Dir, target int, smallerCandidate *int) {
	if currDir.rtotal > target && currDir.rtotal < *smallerCandidate {
		*smallerCandidate = currDir.rtotal
	}
	for dirname := range currDir.dirs {
		collectHigherThan(currDir.dirs[dirname], target, smallerCandidate)
	}

}

func recursiveSum(currDir *Dir, ltotal *int) int {
	currDir.rtotal = sum(*currDir.files)
	for dirname := range currDir.dirs {
		currDir.rtotal += recursiveSum(currDir.dirs[dirname], ltotal)
	}
	if currDir.rtotal < maxDirSize {
		*ltotal += currDir.rtotal
	}
	return currDir.rtotal
}

// Recursiveness here is not needed, but I was deep into it when I noticed.
// Trees and recursion is kind of a reflex.
//
//	In normal situations, I would just start back with something else
func processCommand(lines []string, pos int, currDir *Dir) {
	// circuit-break as below commands call this function
	if pos >= len(lines) {
		return
	}

	if dirname, ok := isCd(lines[pos]); ok {
		if dirname == ".." {
			prevdir := currDir.prev
			processCommand(lines, pos+1, prevdir)
		} else {
			nextdir := currDir.dirs[dirname]
			processCommand(lines, pos+1, nextdir)
		}
	} else {
		if files, dirs, nextPos, ok := isLs(lines, pos, currDir); ok {
			currDir.files = files
			currDir.dirs = dirs
			currDir.total = sum(*files)
			processCommand(lines, nextPos, currDir)
		}
	}
}

func sum(arr []int) (total int) {
	for _, v := range arr {
		total += v
	}
	return
}

func isCd(command string) (string, bool) {
	var dirName string
	if n, _ := fmt.Sscanf(command, "$ cd %s", &dirName); n > 0 {
		return dirName, true
	}
	return "", false
}

func isLs(lines []string, pos int, currDir *Dir) (*[]int, map[string]*Dir, int, bool) {
	if lines[pos] == "$ ls" {
		dirs := map[string]*Dir{}
		files := []int{}
		i := pos + 1
		// first condition below is when last line of inputs file is from ls
		for ; i < len(lines) && !strings.Contains(lines[i], "$"); i++ {
			if strings.HasPrefix(lines[i], "dir") {
				dirName := strings.TrimPrefix(lines[i], "dir ")

				dirs[dirName] = &Dir{prev: currDir}
			} else {
				size := 0
				filename := ""
				fmt.Sscanf(lines[i], "%d %s", &size, &filename)

				files = append(files, size)
			}
		}
		return &files, dirs, i, true
	}
	return nil, nil, 0, false
}
