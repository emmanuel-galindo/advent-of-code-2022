package main

import (
	"bufio"
	"bytes"
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
	inputs, err := os.OpenFile(inputsFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("problem opening %s %v", inputsFilename, err)
	}
	defer inputs.Close()

	data, err := GetData(inputs)
	if err != nil {
		return err
	}
	elves, err := CaloriesPerElf(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	elve := MaxElf(elves)
	fmt.Printf("Elve with most calories is %d with %d calories\n", elve, elves[elve])

	sumTop3Calories := SumTop3(OrderedCalories(elves))
	fmt.Printf("Sum of calories of top 3 elves is %d\n", sumTop3Calories)

	// arr := GetCaloriesPerElf(data)

	return nil
}

func OrderedCalories(calories []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice((calories))))
	return calories
}

func SumTop3(data []int) (sum int) {
	for i := 0; i < 3; i++ {
		sum += data[i]
	}
	return
}

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

func GetData(inputs io.Reader) ([]byte, error) {
	inputsContents, err := io.ReadAll(inputs)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", inputsFilename, err)
	}

	return inputsContents, nil
}
