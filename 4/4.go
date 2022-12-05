package main

import (
	"bufio"
	"examples/adventofcode2022/file_utils"
	"fmt"
	"os"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail       = 1
	inputsFilename = "inputs.txt"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}

}

func run() error {
	inputsFile, closeInputsFile, err := file_utils.InputsFromFile(inputsFilename)
	if err != nil {
		return fmt.Errorf("problem opening %s %v", inputsFilename, err)
	}
	defer closeInputsFile()

	c, o := 0, 0

	scanner := bufio.NewScanner(inputsFile)
	for scanner.Scan() {
		item := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error parsing the file, %v", err)
		}

		if IsContaining(item) {
			c++
		}
		if IsOverlapping(item) {
			o++
		}
	}
	fmt.Printf("%d contained, %d overlapped\n", c, o)
	return nil
}

func IsContaining(item []byte) bool {
	start1, start2, end1, end2 := DecomposeLine(item)
	return (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1)
}
func IsOverlapping(item []byte) bool {
	start1, start2, end1, end2 := DecomposeLine(item)
	return !(end1 < start2 || end2 < start1)
}

func DecomposeLine(item []byte) (start1 int, start2 int, end1 int, end2 int) {
	fmt.Sscanf(string(item), "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
	return
}
