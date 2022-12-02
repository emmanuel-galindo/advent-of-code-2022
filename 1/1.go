package main

import (
	"bufio"
	"examples/adventofcode2022/file_utils"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
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

	elves, err := CaloriesPerElf(inputsFile)
	if err != nil {
		return fmt.Errorf("problem processing calories data from %s %v", inputsFilename, err)
	}
	elve := MaxElf(elves)
	fmt.Printf("Elve with most calories is %d with %d calories\n", elve, elves[elve])

	sumTop3Calories := SumTop3(OrderedCalories(elves))
	fmt.Printf("Sum of calories of top 3 elves is %d\n", sumTop3Calories)

	return nil
}

// InputsFromFile opens a file in the filesystem and returns a File pointer
// OrderedCalories sorts the input calories arrays in descending order
func OrderedCalories(calories []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice((calories))))
	return calories
}

// SumTop3 returns the sum of the calories from the three top elves
func SumTop3(data []int) (sum int) {
	for i := 0; i < 3; i++ {
		sum += data[i]
	}
	return
}

// CaloriesPerElf processes the input data and returns calories grouped by elf
func CaloriesPerElf(data io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(data)
	var (
		elves []int
		sum   int
	)
	for scanner.Scan() {
		switch line := scanner.Text(); line {
		case "":
			elves = append(elves, sum)
			sum = 0
		default:
			num, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			sum += num
		}
	}
	if sum > 0 {
		elves = append(elves, sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return elves, nil
}

// MaxElf returns the elf that collected more calories
func MaxElf(elves []int) int {
	var max, maxi int
	for i, v := range elves {
		if v > max {
			max = v
			maxi = i
		}
	}
	return maxi
}
