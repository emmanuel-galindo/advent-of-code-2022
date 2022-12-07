package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
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

func run() error {
	inputsArr := strings.Split(inputsContent, "\n\n")
	stacksInput := strings.Split(inputsArr[0], "\n")
	operationsInput := strings.Split(inputsArr[1], "\n")

	// Get the quantity of stacks from the last line
	lastLine := stacksInput[len(stacksInput)-1:][0]
	qty_stacks, _ := strconv.Atoi(string(lastLine)[len(lastLine)-2 : len(lastLine)-1])

	// Remove last line
	stacksInput = stacksInput[:len(stacksInput)-1]

	stacks := ParseStacks(stacksInput, qty_stacks)
	stacksMultiple := ParseStacks(stacksInput, qty_stacks)

	// Execute operations one crate at a time
	asStack(stacks, operationsInput)
	// Move all specified crates at once
	asMultiple(stacksMultiple, operationsInput)

	fmt.Printf("CrateMover 9000 ordered as %s\n", GetTotal(stacks))
	fmt.Printf("CrateMover 9001 ordered as %s\n", GetTotal(stacksMultiple))

	return nil
}

func ParseStacks(stacksInput []string, qty_stacks int) [][]string {
	stacks := make([][]string, qty_stacks)
	for i := len(stacksInput) - 1; i >= 0; i-- {
		v := stacksInput[i]
		idx := 2

		for i := 0; i < qty_stacks; i++ {
			crate := string(v)[idx-1 : idx]
			idx += 4

			if crate != " " {
				stacks[i] = append(stacks[i], crate)
			}
		}
	}
	return stacks
}

func asMultiple(stacksMultiple [][]string, operationsInput []string) {
	ExecuteOperations(stacksMultiple, operationsInput, true)
}

func asStack(stacks [][]string, operationsInput []string) {
	ExecuteOperations(stacks, operationsInput, false)
}

func ExecuteOperations(stacks [][]string, operationsInput []string, multiple bool) [][]string {
	for _, operation := range operationsInput {
		var qty_crates, src, dst int

		fmt.Sscanf(operation, "move %d from %d to %d", &qty_crates, &src, &dst)
		src_stack := stacks[src-1]
		dst_stack := stacks[dst-1]

		if multiple {
			start := len(src_stack) - qty_crates

			// pop
			crates := src_stack[start:]
			src_stack = src_stack[:start]

			dst_stack = append(dst_stack, crates...)
		} else {
			for i := 0; i < qty_crates; i++ {
				n := len(src_stack) - 1

				// pop
				crate := src_stack[n]
				src_stack = src_stack[:n]

				dst_stack = append(dst_stack, crate)
			}
		}

		stacks[src-1] = src_stack
		stacks[dst-1] = dst_stack
	}
	return stacks
}

func GetTotal(stacks [][]string) (c string) {
	for _, v := range stacks {
		c += v[len(v)-1]
	}
	return
}
