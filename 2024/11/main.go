package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

var transformCache = make(map[int][]int)

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

	nums := prepareInput(input)
	for range 25 {
		var newNums []int
		for _, num := range nums {
			newNums = append(newNums, transformNumber(num)...)
		}
		nums = newNums
	}
	fmt.Println(len(nums))
}

func secondPart() {
	slog.Debug("Running second part")
}

func prepareInput(input string) []int {
	parts := strings.Fields(input)
	var numbers []int
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("can't convert %v to int", part))
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func transformNumber(n int) []int {
	cachedTransform, ok := transformCache[n]
	var result []int
	if ok {
		return cachedTransform
	}

	digitAmount := getDigitAmount(n)
	switch {
	case n == 0:
		result = []int{1}
	case digitAmount%2 == 0:
		result = []int{n / int(math.Pow10(digitAmount/2)), n % int(math.Pow10(digitAmount/2))}
	default:
		result = []int{n * 2024}
	}
	transformCache[n] = result

	return result
}

func getDigitAmount(n int) int {
	if n == 0 {
		return 1
	}
	amount := 0
	for n != 0 {
		n /= 10
		amount++
	}
	return amount
}
