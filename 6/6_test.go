package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed inputs.txt
var inputsTest string

func BenchmarkStartOffset(b *testing.B) {
	lines := strings.Split(inputsTest, "\n")

	b.ResetTimer()
	for _, line := range lines {
		fmt.Printf("StartOffset: 4 chars => %d, 14 chars => %d\n",
			nestedLoops(line, 4), nestedLoops(line, 14))
	}
}

func BenchmarkOnePass(b *testing.B) {
	lines := strings.Split(inputsTest, "\n")
	b.ResetTimer()
	for _, line := range lines {
		fmt.Printf("OnePass: 4 chars => %d, 14 chars => %d\n",
			slider(line, 4), slider(line, 14))
	}
}
