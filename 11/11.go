package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var inputsContent string

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

const (
	lineNumber = iota
	lineItems
	lineOperation
	lineTest
	lineTestTrue
	lineTestFalse
)

type Test struct {
	divisibleBy int
	op          func(int) int
}

type Monkey struct {
	number    int
	items     []int
	operation func(int) int
	test      Test
	inspected int
	rawData   []string
	// test
}

func NewMonkey(rawData []string) Monkey {
	monkey := Monkey{rawData: rawData}
	monkey.ParseNumber()
	monkey.ParseItems()
	monkey.ParseOperation()
	monkey.ParseTest()

	return monkey
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %v; inspected %d ", m.number, m.items, m.inspected)
}

func (m *Monkey) ParseNumber() {
	fmt.Sscanf(m.rawData[lineNumber], "Monkey %d:", &m.number)
}

func (m *Monkey) ParseItems() {
	m.items = strArrToInt(strings.Split(strings.TrimPrefix(m.rawData[lineItems], "  Starting items: "), ", "))
}

func (m *Monkey) ParseOperation() {
	str := strings.TrimPrefix(m.rawData[lineOperation], "  Operation: new = ")
	var modifier int
	switch {
	case str == "old * old":
		m.operation = func(old int) int {
			return old * old
		}
	case strings.HasPrefix(str, "old * "):
		fmt.Sscanf(str, "old * %d", &modifier)
		m.operation = func(old int) int {
			return old * modifier
		}
	case strings.HasPrefix(str, "old + "):
		fmt.Sscanf(str, "old + %d", &modifier)
		m.operation = func(old int) int {
			return old + modifier
		}
	}
}

func (m *Monkey) ParseTest() {
	testLines := m.rawData[lineTest:]

	// todo:create Test and send the property
	var divisible, whenTrue, whenFalse int
	fmt.Sscanf(testLines[0], "  Test: divisible by %d", &divisible)
	fmt.Sscanf(testLines[1], "    If true: throw to monkey %d", &whenTrue)
	fmt.Sscanf(testLines[2], "    If false: throw to monkey %d", &whenFalse)

	m.test = Test{
		divisibleBy: divisible,
		op: func(item int) int {
			if item%divisible == 0 {
				return whenTrue
			}
			return whenFalse
		},
	}
}

type Processor struct {
	bigMod      int
	monkeys     map[int]*Monkey
	monkeysData []string
}

func NewProcessor() Processor {
	return Processor{}
}

func (p *Processor) ParseMonkeysData(data string) {
	p.monkeysData = strings.Split(inputsContent, "\n\n")
}

func (p *Processor) CreateMonkeys() {
	monkeys := map[int]*Monkey{}
	for _, monkeyData := range p.monkeysData {
		monkey := p.CreateMonkey(monkeyData)
		monkeys[monkey.number] = &monkey
	}
	p.monkeys = monkeys
}

func (p *Processor) CreateMonkey(monkeyData string) Monkey {
	lines := strings.Split(monkeyData, "\n")
	return NewMonkey(lines)
}

func (p *Processor) ExecuteRounds(rounds, divisible int) {
	for r := 0; r < rounds; r++ {
		for i := 0; i < len(p.monkeys); i++ {
			m := p.monkeys[i]
			for _, item := range m.items {
				m.inspected++

				// lets try to "keep worry level manageable"
				if divisible == 1 {
					smallMod := item % p.bigMod
					item = smallMod + p.bigMod
				}
				item = m.operation(item) / divisible

				next := m.test.op(item)
				p.monkeys[next].items = append(p.monkeys[next].items, item)

				m.items = shift(m.items)
			}
		}
	}
}

func (p *Processor) CalculateMonkeyBusiness() int {
	inspected := []int{}
	for i := 0; i < len(p.monkeys); i++ {
		inspected = append(inspected, p.monkeys[i].inspected)
	}

	sort.Slice(inspected, func(i, j int) bool { return inspected[i] > inspected[j] })
	return inspected[0] * inspected[1]

}

func (p *Processor) CalculateBigMod() {
	var bigMod int = 1
	for _, monkey := range p.monkeys {
		bigMod *= monkey.test.divisibleBy
	}
	p.bigMod = bigMod
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run() error {
	processor := NewProcessor()
	processor.ParseMonkeysData(inputsContent)
	processor.CreateMonkeys()

	processor.ExecuteRounds(20, 3.0)
	result := processor.CalculateMonkeyBusiness()
	fmt.Printf("Result part1: %d\n", result)

	processor.CreateMonkeys() // reset
	processor.CalculateBigMod()
	processor.ExecuteRounds(10000, 1.0)
	result = processor.CalculateMonkeyBusiness()
	fmt.Printf("Result part2: %d\n", result)

	return nil
}

// Helpers
func shift(arr []int) []int {
	if len(arr) > 1 {
		return arr[1:]
	}
	return []int{}
}

func strArrToInt(strArr []string) (intArr []int) {
	for _, str := range strArr {
		number, _ := strconv.Atoi(str)
		intArr = append(intArr, int(number))
	}
	return
}
