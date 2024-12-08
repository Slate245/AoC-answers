package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type calibrationEntry struct {
	testValue int
	operands  []int
}

var operators = []func(int, int) int{func(a int, b int) int { return a + b }, func(a int, b int) int { return a * b }}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	shouldRunSecondPart := flag.Bool("part2", false, "second part solution")
	flag.Parse()

	if shouldRunSecondPart != nil && *shouldRunSecondPart {
		secondPart()
		return
	}

	firstPart()
}

func firstPart() {
	slog.Debug("Running first part")

	entries := prepareInput(input)
	result := 0
	for _, entry := range entries {
		isPossible := false
		possibleResults := []int{entry.operands[0]}
		for i := 1; i < len(entry.operands); i++ {
			possibleResults = applyOperators(possibleResults, entry.operands[i], entry.testValue)
			if len(possibleResults) == 0 {
				break
			}
		}
		for _, pr := range possibleResults {
			if pr == entry.testValue {
				isPossible = true
				break
			}
		}
		if isPossible {
			result += entry.testValue
		}
	}

	fmt.Println(result)
}

func secondPart() {
	slog.Debug("Running second part")
	operators = append(operators, func(i1, i2 int) int {
		s1 := strconv.Itoa(i1)
		s2 := strconv.Itoa(i2)

		var b strings.Builder
		b.WriteString(s1)
		b.WriteString(s2)

		i, err := strconv.Atoi(b.String())
		if err != nil {
			panic(fmt.Sprintf("Can't convert %v to int", i))
		}

		return i
	})

	entries := prepareInput(input)
	result := 0
	for _, entry := range entries {
		isPossible := false
		possibleResults := []int{entry.operands[0]}
		for i := 1; i < len(entry.operands); i++ {
			possibleResults = applyOperators(possibleResults, entry.operands[i], entry.testValue)
			if len(possibleResults) == 0 {
				break
			}
		}
		for _, pr := range possibleResults {
			if pr == entry.testValue {
				isPossible = true
				break
			}
		}
		if isPossible {
			result += entry.testValue
		}
	}

	fmt.Println(result)

}

func applyOperators(possibleResults []int, nextValue int, checkValue int) []int {
	var results []int
	for _, operator := range operators {
		for _, pr := range possibleResults {
			nextResult := operator(pr, nextValue)
			if nextResult > checkValue {
				continue
			}
			results = append(results, nextResult)
		}
	}

	return results
}

func prepareInput(input string) []calibrationEntry {
	lines := strings.Split(input, "\n")
	preparedInput := make([]calibrationEntry, len(lines))
	for i, line := range lines {
		var entry calibrationEntry
		parts := strings.Split(line, ":")
		tv, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("Can't convert %v to int", parts[0]))
		}
		strOperands := strings.Fields(parts[1])
		operands := make([]int, len(strOperands))
		for i, so := range strOperands {
			o, err := strconv.Atoi(so)
			if err != nil {
				panic(fmt.Sprintf("Can't convert %v to int", so))
			}
			operands[i] = o
		}
		entry.testValue = tv
		entry.operands = operands
		preparedInput[i] = entry
	}

	return preparedInput
}
