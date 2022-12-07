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

	for _, line := range lines {
		fmt.Printf("StartOffset: 4 chars => %d, 14 chars => %d\n",
			getStartOffset(line, 4), getStartOffset(line, 14))
		fmt.Printf("OnePass: 4 chars => %d, 14 chars => %d\n",
			OnePass(line, 4), OnePass(line, 14))
	}
}

func OnePass(input string, diff int) int {
	seen := map[rune]int{}
	c := 0
	line := string(input)
	for idx := 0; idx < len(line); idx++ {
		char := rune(line[idx])

		if _, present := seen[char]; present {
			idx = seen[char]
			seen = map[rune]int{}
			c = 0
			continue
		}

		c += 1
		if c == diff {
			return idx + 1
		}

		seen[char] = idx
	}
	return 0
}

func getStartOffset(line string, diff int) int {
	lineArr := string(line)
	for i := 0; i < len(lineArr); i++ {
		substr := line[i : i+diff]
		if differentChars(substr) {
			return i + diff
		}
	}
	return 0

}

func differentChars(str string) bool {
	seen := map[rune]bool{}
	for _, char := range string(str) {
		// if seen[char] {
		// return false
		// }
		seen[char] = true
	}
	// return true
	return (len(seen) == len(str))
}
