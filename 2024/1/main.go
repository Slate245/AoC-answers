package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	shouldRunSecondPart := flag.Bool("part2", false, "second part solution")
	flag.Parse()

	if shouldRunSecondPart != nil && *shouldRunSecondPart {
		secondPart()
		return
	}

	firstPart()
}

func firstPart() {
	var firstList []int
	var secondList []int

	for _, line := range strings.Split(input, "\n") {
		ids := strings.Fields(line)
		firstId, err1 := strconv.Atoi(ids[0])
		secondId, err2 := strconv.Atoi(ids[1])

		if err1 != nil || err2 != nil {
			panic("incorrect input")
		}

		firstList = append(firstList, int(firstId))
		secondList = append(secondList, int(secondId))
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	var sum int
	for i := 0; i < len(firstList); i++ {
		diff := firstList[i] - secondList[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	fmt.Print(sum)
}

func secondPart() {
	var firstList []int
	var secondList []int

	for _, line := range strings.Split(input, "\n") {
		ids := strings.Fields(line)
		firstId, err1 := strconv.Atoi(ids[0])
		secondId, err2 := strconv.Atoi(ids[1])

		if err1 != nil || err2 != nil {
			panic("incorrect input")
		}

		firstList = append(firstList, int(firstId))
		secondList = append(secondList, int(secondId))
	}

	occurenceNumberMap := make(map[int]int, len(firstList))
	for _, num := range secondList {
		occurenceNumberMap[num] += 1
	}

	var similarityScore int
	for _, num := range firstList {
		similarityScore += num * occurenceNumberMap[num]
	}

	fmt.Print(similarityScore)
}
