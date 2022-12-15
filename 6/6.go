package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed inputs.txt
var inputsContent string

func main() {
	lines := strings.Split(inputsContent, "\n")

	for _, line := range lines {

		start := time.Now()
		fmt.Printf("nested loops: 4 chars => %d, 14 chars => %d\n",
			nestedLoops(line, 4), nestedLoops(line, 14))
		elapsed1 := time.Since(start) / 10000

		start = time.Now()
		fmt.Printf("slider: 4 chars => %d, 14 chars => %d\n",
			slider(line, 4), slider(line, 14))
		elapsed2 := time.Since(start) / 10000

		start = time.Now()
		fmt.Printf("hashSize: 4 chars => %d, 14 chars => %d\n",
			hashSize(line, 4), hashSize(line, 14))
		elapsed3 := time.Since(start) / 10000

		start = time.Now()
		fmt.Printf("bitmask: 4 chars => %d, 14 chars => %d\n",
			bitmaskOR(line, 4), bitmaskOR(line, 14))
		elapsed4 := time.Since(start) / 10000

		start = time.Now()
		fmt.Printf("bitmask: 4 chars => %d, 14 chars => %d\n",
			bitmaskXOR(line, 4), bitmaskXOR(line, 14))
		elapsed5 := time.Since(start) / 10000

		fmt.Printf("nested loops%s\n", elapsed1)
		fmt.Printf("slider %s\n", elapsed2)
		fmt.Printf("hashSize %s\n", elapsed3)
		fmt.Printf("bitmask OR %s\n", elapsed4)
		fmt.Printf("bitmask XOR %s\n", elapsed5)
	}
}

func bitmaskOR(line string, diff int) int {
	charsArr := string(line)
	for i := 0; i < len(charsArr)-diff; i++ {
		var set uint32
		for j := 0; j < diff; j++ {
			set |= 1 << (int32(line[i+j] - 'a'))
		}
		bitmaskStr := fmt.Sprintf("%08b\n", set)
		if strings.Count(bitmaskStr, "1") == diff {
			return i + diff
		}
	}
	return 0
}

// https://www.mattkeeter.com/blog/2022-12-10-xor/
func bitmaskXOR(line string, diff int) int {
	charsArr := string(line)
	var set uint32
	for i := 0; i < len(charsArr)-diff; i++ {
		set ^= 1 << (int32(line[i] - 'a'))

		if i >= diff {
			set ^= 1 << (int32(line[i-diff] - 'a'))
		}
		bitmaskStr := fmt.Sprintf("%08b\n", set)
		if strings.Count(bitmaskStr, "1") == diff {
			return i + diff
		}
	}
	return 0
}

func hashSize(line string, diff int) int {
	charsArr := string(line)
	for i := 0; i < len(charsArr)-diff; i++ {
		seen := map[rune]bool{}
		for j := 0; j < diff; j++ {
			seen[rune(line[i+j])] = true
		}

		if len(seen) == diff {
			return i + diff
		}
	}
	return 0

}

func slider(input string, diff int) int {
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

func nestedLoops(line string, diff int) int {
	charsArr := string(line)
	for i := 0; i < len(charsArr)-diff; i++ {
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
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
