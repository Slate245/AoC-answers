package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"strings"
)

type Direction int

//go:embed input.txt
var input string

var SEARCH_TERM = []string{"X", "M", "A", "S"}

const (
	Forward Direction = iota
	Backward
	Upward
	Downward
	UpForward
	UpBackward
	DownForward
	DownBackward
)

var directionNames = map[Direction]string{
	Forward:      "forward",
	Backward:     "backward",
	Upward:       "upward",
	Downward:     "downward",
	UpForward:    "upward and forward",
	UpBackward:   "upward and backward",
	DownForward:  "downward and forward",
	DownBackward: "downward and backward",
}

var increments = map[Direction][2]int{
	Forward:      {0, 1},
	Backward:     {0, -1},
	Upward:       {-1, 0},
	Downward:     {1, 0},
	UpForward:    {-1, 1},
	UpBackward:   {-1, -1},
	DownForward:  {1, 1},
	DownBackward: {1, -1},
}

func (d Direction) String() string {
	return directionNames[d]
}

func (d Direction) CheckSpace(input [][]string, index [2]int) bool {
	switch d {
	case Forward:
		slog.Debug(fmt.Sprintf("Checking available space forward at %v", index))
		return index[1] <= len(input[index[0]])-len(SEARCH_TERM)
	case Backward:
		slog.Debug(fmt.Sprintf("Checking available space backward at %v", index))
		return index[1] >= len(SEARCH_TERM)-1
	case Upward:
		slog.Debug(fmt.Sprintf("Checking available space upward at %v", index))
		return index[0] >= len(SEARCH_TERM)-1
	case Downward:
		slog.Debug(fmt.Sprintf("Checking available space downward at %v", index))
		return index[0] <= len(input)-len(SEARCH_TERM)
	case UpForward:
		return Upward.CheckSpace(input, index) && Forward.CheckSpace(input, index)
	case UpBackward:
		return Upward.CheckSpace(input, index) && Backward.CheckSpace(input, index)
	case DownForward:
		return Downward.CheckSpace(input, index) && Forward.CheckSpace(input, index)
	case DownBackward:
		return Downward.CheckSpace(input, index) && Backward.CheckSpace(input, index)
	default:
		return false
	}
}

func (d Direction) GetIncrement() [2]int {
	return increments[d]
}

var directions = []Direction{Upward, UpForward, Forward, DownForward, Downward, DownBackward, Backward, UpBackward}

func (d Direction) Search(input [][]string, index [2]int) bool {
	slog.Debug(fmt.Sprintf("Searching %v", d))
	if !d.CheckSpace(input, index) {
		return false
	}
	increments := d.GetIncrement()
	for i, str := range SEARCH_TERM {
		strToCompare := input[index[0]+i*increments[0]][index[1]+i*increments[1]]
		slog.Debug(fmt.Sprintf("Comparing %s to %s", strToCompare, str))
		if strToCompare != str {
			slog.Debug("Word not found")
			return false
		}
	}

	slog.Debug("Found word")
	return true
}

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
	searchInput := prepareInput(input)

	var numberOfWords int

	for i := 0; i < len(searchInput); i++ {
		for j := 0; j < len(searchInput[i]); j++ {
			for _, dir := range directions {
				if dir.Search(searchInput, [2]int{i, j}) {
					numberOfWords += 1
				}
			}
		}
	}

	fmt.Printf("Found %d words\n", numberOfWords)
}

func secondPart() {
	slog.Debug("Running second part")

}

func prepareInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	preparedInput := make([][]string, len(lines))
	for i, line := range lines {
		preparedInput[i] = strings.Split(line, "")
	}

	return preparedInput
}
