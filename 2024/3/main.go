package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var MUL_INSTRUCTION_REGEXP = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
var DISABLET_INSTRUCTIONS_REGEXP = regexp.MustCompile(`don't\(\)(.|[\s\S])*?(do\(\)|$)`)

func main() {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
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
	result := computeResult(input)

	fmt.Printf("Final result: %d\n", result)
}

func computeResult(inputStr string) int {
	instructions := MUL_INSTRUCTION_REGEXP.FindAllString(inputStr, -1)
	slog.Info(fmt.Sprintf("Found %d mul instructions", len(instructions)))
	var result int

	for _, instruction := range instructions {
		slog.Debug("executing instruction", slog.String("instruction", instruction))
		strArgs := strings.Split(strings.Trim(instruction, "mul()"), ",")
		arguments := make([]int, 2)
		for i, strArg := range strArgs {
			arg, err := strconv.Atoi(strArg)

			if err != nil {
				panic(fmt.Sprintf("Unable to parse %v as int", strArg))
			}
			arguments[i] = arg
		}
		result += arguments[0] * arguments[1]
		slog.Debug("Updated result", slog.Int("result", result))
	}
	return result
}

func secondPart() {
	slog.Debug("Running second part")
	instructionsWithoutDisabled := DISABLET_INSTRUCTIONS_REGEXP.ReplaceAllString(input, "")

	result := computeResult(instructionsWithoutDisabled)

	fmt.Printf("Final result: %d\n", result)
}
