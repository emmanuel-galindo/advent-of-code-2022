package main

import (
	"bufio"
	"examples/adventofcode2022/file_utils"
	"fmt"
	"io"
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

	total, totalGroup, err := ProcessItems(inputsFile)
	if err != nil {
		return err
	}
	fmt.Printf("The sum of all priorities is %d, and groups %d", total, totalGroup)

	return nil
}

func SplitItems(line string) (firstHalf, secondHalf []byte) {
	data_arr := []byte(line)
	cut := len(data_arr) / 2
	firstHalf = data_arr[0:cut]
	secondHalf = data_arr[cut:]
	return
}

func FindInterjectItems(items1, items2 []byte) (interjectedItems []byte) {

	items := map[byte]uint8{}
	for _, v := range items1 {
		items[v] |= (1 << 0)
	}
	for _, v := range items2 {
		items[v] |= (1 << 1)
	}

	// var interjectedItems, inAButNotB, inBButNotA []int
	for k, v := range items {
		item1 := v&(1<<0) != 0
		item2 := v&(1<<1) != 0
		switch {
		case item1 && item2:
			interjectedItems = append(interjectedItems, k)
		}
	}
	return
}

// PrioritizeItems checks each element from the provided list of chars
// and returns the list with the equivalent priority from its ascii value
// a to z; 1 to 26
// A to Z; 27 to 52
func PrioritizeItems(data []byte) ([]int, error) {
	items := []int{}
	for _, v := range data {
		// in ascii table,
		// lowercases are from 97 to 122
		// uppercases are from 65 to 90
		// equivalences are 1-based so it has to be compensated
		if v >= 97 && v <= 122 {
			items = append(items, int(v)%96)
		} else if v >= 65 && v <= 90 {
			items = append(items, int(v)%64+26)
		} else {
			return nil, fmt.Errorf("item %s is not valid", string(v))
		}
	}
	return items, nil
}

func SumItems(input []int) (total int) {
	for i := range input {
		total += input[i]
	}
	return
}

func ProcessItems(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)
	interjects, groupInterject := []byte{}, []byte{}
	groupOf3 := []string{}
	c := 1

	for scanner.Scan() {
		item := scanner.Text()
		if err := scanner.Err(); err != nil {
			return 0, 0, fmt.Errorf("error parsing the file, %v", err)
		}

		interjects = append(interjects, FindInterjectItems(SplitItems(item))...)

		groupOf3 = append(groupOf3, item)
		if c%3 == 0 {
			arr1, arr2 := []byte(groupOf3[0]), []byte(groupOf3[1])
			groupInterject = append(groupInterject, FindInterjectItems(FindInterjectItems(arr1, arr2), []byte(groupOf3[2]))...)
			groupOf3 = groupOf3[:0]
		}
		c += 1
	}
	prioritizedItems, err := PrioritizeItems(interjects)
	if err != nil {
		return 0, 0, fmt.Errorf("error prioriting items, %v", err)
	}
	groupPrioritizedItems, err := PrioritizeItems(groupInterject)
	if err != nil {
		return 0, 0, fmt.Errorf("error prioriting items from  group, %v", err)
	}

	return SumItems(prioritizedItems), SumItems(groupPrioritizedItems), nil
}
