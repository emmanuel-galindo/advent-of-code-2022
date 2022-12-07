package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

type Dir struct {
	dirs   map[string]*Dir
	files  *[]int
	prev   *Dir
	total  int
	rtotal int
}

func main() {
	lines := strings.Split(inputsContent, "\n")

	// TODO send pointers instead of DIR
	root := &Dir{prev: nil}
	processCommand(lines, 1, root)
	var total int = 0
	RecursiveSum(root, &total)
	// ltotal := SumLowestThan(root, 100000)
	fmt.Printf("sum of dirs with recursive total less than < 100000 => %d\n", total)

	target := 30000000 - (70000000 - root.rtotal)
	fmt.Printf("target => %d\n", target)
	smallerCandidate := root.rtotal
	collectHigherThan(root, target, &smallerCandidate)
	fmt.Printf("%d", smallerCandidate)
}

func collectHigherThan(currDir *Dir, target int, smallerCandidate *int) {
	if currDir.rtotal > target && currDir.rtotal < *smallerCandidate {
		*smallerCandidate = currDir.rtotal
	}
	for dirname := range currDir.dirs {
		collectHigherThan(currDir.dirs[dirname], target, smallerCandidate)
	}

}

func RecursiveSum(currDir *Dir, ltotal *int) int {
	sumTotal := sum(*currDir.files)
	currDir.rtotal = sumTotal
	for dirname := range currDir.dirs {
		currDir.rtotal += RecursiveSum(currDir.dirs[dirname], ltotal)
	}
	if currDir.rtotal < 100000 {
		*ltotal += currDir.rtotal
	}
	return currDir.rtotal
}

func processCommand(lines []string, pos int, currDir *Dir) {
	if pos >= len(lines) {
		return
	}
	if dirname, ok := isCd(lines[pos]); ok {
		// assuming ls will always come before cd
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
			totalSize := sum(*files)
			currDir.total = totalSize
			// if currDir.prev != nil {
			// currDir.prev.total += totalSize
			// }
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
	var dirname string
	if n, _ := fmt.Sscanf(command, "$ cd %s", &dirname); n > 0 {
		return dirname, true
	}
	return "", false
}

func isLs(lines []string, pos int, currDir *Dir) (*[]int, map[string]*Dir, int, bool) {
	if lines[pos] == "$ ls" {
		dirs := map[string]*Dir{}
		files := []int{}
		i := pos + 1
		for ; i < len(lines) && !strings.Contains(lines[i], "$"); i++ {
			if strings.HasPrefix(lines[i], "dir") {
				dirName := ""
				fmt.Sscanf(lines[i], "dir %s", &dirName)
				newDir := &Dir{prev: currDir}

				dirs[dirName] = newDir
			} else {
				size := 0
				filename := ""
				// TODO: filename is not used, only for sscanf to compile
				fmt.Sscanf(lines[i], "%d %s", &size, &filename)

				files = append(files, size)
			}
		}
		return &files, dirs, i, true
	}
	return nil, nil, 0, false
}
