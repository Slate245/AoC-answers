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

var machineRulesRegexp = regexp.MustCompile("Button A: (?<a>.*\n)Button B: (?<b>.*\n)Prize: (?<prize>.*)")
var buttonRegexp = regexp.MustCompile(`[+-]\d+`)
var prizeRegexp = regexp.MustCompile(`\d+`)

type MachineRules struct {
	prizeLocation struct {
		x int
		y int
	}
	aButtonRules struct {
		x int
		y int
	}
	bButtonRules struct {
		x int
		y int
	}
}

func (mr *MachineRules) GetLowestTokenAmountForWin() int {
	if (mr.aButtonRules.x*mr.prizeLocation.y-mr.aButtonRules.y*mr.prizeLocation.x)%(mr.aButtonRules.x*mr.bButtonRules.y-mr.bButtonRules.x*mr.aButtonRules.y) != 0 {
		return 0
	}
	bButtonPresses := (mr.aButtonRules.x*mr.prizeLocation.y - mr.aButtonRules.y*mr.prizeLocation.x) / (mr.aButtonRules.x*mr.bButtonRules.y - mr.bButtonRules.x*mr.aButtonRules.y)
	if (mr.prizeLocation.x-bButtonPresses*mr.bButtonRules.x)%mr.aButtonRules.x != 0 {
		return 0
	}
	aButtonPresses := (mr.prizeLocation.x - bButtonPresses*mr.bButtonRules.x) / mr.aButtonRules.x

	if aButtonPresses > 100 || bButtonPresses > 100 {
		return 0
	}
	const aButtonCost = 3
	const bButtonCost = 1
	return aButtonPresses*aButtonCost + bButtonPresses*bButtonCost
}

func newMachineRules(s string) *MachineRules {
	mr := MachineRules{}

	matches := machineRulesRegexp.FindStringSubmatch(s)[1:]

	for i := range 2 {
		numStrs := buttonRegexp.FindAllString(matches[i], -1)
		var nums []int
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to int", numStr))
			}
			nums = append(nums, num)
		}

		if i == 0 {
			mr.aButtonRules.x = nums[0]
			mr.aButtonRules.y = nums[1]
		} else {
			mr.bButtonRules.x = nums[0]
			mr.bButtonRules.y = nums[1]
		}
	}

	numStrs := prizeRegexp.FindAllString(matches[2], -1)
	var nums []int
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(fmt.Sprintf("can't convert %v to int", numStr))
		}
		nums = append(nums, num)
	}
	mr.prizeLocation.x = nums[0]
	mr.prizeLocation.y = nums[1]

	return &mr
}

//go:embed input.txt
var input string

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

	var machines []*MachineRules

	preparedInput := preparedInput(input)
	for _, lines := range preparedInput {
		mr := newMachineRules(lines)
		machines = append(machines, mr)
	}

	tokenAmount := 0

	for _, mr := range machines {
		tokenAmount += mr.GetLowestTokenAmountForWin()
	}

	fmt.Println(tokenAmount)
}

func secondPart() {
	slog.Debug("Running second part")
}

func preparedInput(input string) []string {
	return strings.Split(input, "\n\n")
}
