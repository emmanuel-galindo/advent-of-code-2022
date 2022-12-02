package main

import (
	"bufio"
	"examples/adventofcode2022/file_utils"
	"fmt"
	"io"
	"os"
	"strings"
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

	rounds, err := GetRounds(inputsFile)
	if err != nil {
		return err
	}
	totalScore, totalScoreAfterKnowingTheSecret := ProcessRounds(rounds)

	fmt.Printf("My total score is %d, but after the elves told me the rest of the secret it is %d", totalScore, totalScoreAfterKnowingTheSecret)

	return nil
}

func GetRounds(input io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(input)
	var (
		rounds [][]string
	)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		rounds = append(rounds, arr)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing the file, %v", err)
	}
	return rounds, nil
}

func ProcessRound(round []string) (int, int) {
	// Rock (A,X) > Scissors (C,Z)
	// Paper (B,Y) > Rock (A,X)
	// Scissors (C,Z) > Paper (B,Y)
	var (
		myHand       string = round[1]
		opponentHand string = round[0]
	)
	// behavior := map[string]map[string]string{
	// 	"X": {
	// 		"win":  "C",
	// 		"lose": "B",
	// 		"draw": "A",
	// 	},
	// 	"Y": {
	// 		"win":  "A",
	// 		"lose": "C",
	// 		"draw": "B",
	// 	},
	// 	"Z": {
	// 		"win":  "B",
	// 		"lose": "A",
	// 		"draw": "Z",
	// 	},
	// }

	wins := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	equals := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	lose := map[string]string{
		"X": "B",
		"Y": "C",
		"Z": "A",
	}
	score := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	secret := func() int {
		if wins[myHand] == opponentHand {
			return 6 + score[myHand]
		}
		if equals[myHand] == opponentHand {
			return 3 + score[myHand]
		}
		return score[myHand]
	}
	scoreSecret := secret()

	completeSecret := map[string]func() int{
		"X": func() int { return score[findByValue(lose, opponentHand)] },
		"Y": func() int { return 3 + score[findByValue(equals, opponentHand)] },
		"Z": func() int { return 6 + score[findByValue(wins, opponentHand)] },
	}
	scoreCompleteSecret := completeSecret[myHand]()

	return scoreSecret, scoreCompleteSecret

	// if wins[myHand] == opponentHand {
	// 	return 6 + score[myHand], score2
	// }
	// if equals[myHand] == opponentHand {
	// 	return 3 + score[myHand], score2
	// }
	// return score[myHand], score2
}

func findByValue(arr map[string]string, t string) string {
	for k, v := range arr {
		if v == t {
			return k
		}
	}
	return ""
}

func ProcessRounds(rounds [][]string) (total, total2 int) {
	for _, round := range rounds {
		sc1, sc2 := ProcessRound(round)
		total += sc1
		total2 += sc2
	}
	return
}
