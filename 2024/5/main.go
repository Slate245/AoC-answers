package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

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

	rules, updates := prepareInput(input)
	var validUpdates [][]string

	for _, update := range updates {
		slog.Debug(fmt.Sprintf("Checking update: %v", update))
		isValid := true
		for _, rule := range rules {
			slog.Debug(fmt.Sprintf("Checking rule %v", rule))

			firstIndex := slices.Index(update, rule[0])
			secondIndex := slices.Index(update, rule[1])
			slog.Debug(fmt.Sprintf("Index of %s is %d, index of %s is %d", rule[0], firstIndex, rule[1], secondIndex))
			if firstIndex == -1 || secondIndex == -1 {
				slog.Debug(fmt.Sprintf("Rule %v unapplicable", rule))
				continue
			}
			if firstIndex > secondIndex {
				slog.Debug(fmt.Sprintf("Rule %v is broken", rule))
				isValid = false
				break
			}
		}
		if isValid {
			validUpdates = append(validUpdates, update)
		}
	}
	slog.Debug(fmt.Sprintf("Valid updates: %v", validUpdates))
	slog.Debug(fmt.Sprintf("Found %d valid updates", len(validUpdates)))

	result := 0
	for _, update := range validUpdates {
		middleNumber, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		slog.Debug(fmt.Sprintf("Middle number in %v is %d", update, middleNumber))
		result += middleNumber
	}
	fmt.Println(result)
}

func secondPart() {
	slog.Debug("Running second part")
}

func prepareInput(input string) ([][2]string, [][]string) {
	parts := strings.Split(input, "\n\n")
	rulesStrings := strings.Split(parts[0], "\n")
	updatesStrings := strings.Split(parts[1], "\n")

	rules := make([][2]string, len(rulesStrings))
	for i, line := range rulesStrings {
		rule := strings.Split(line, "|")
		rules[i] = [2]string{rule[0], rule[1]}
	}
	updates := make([][]string, len(updatesStrings))
	for i, line := range updatesStrings {
		updates[i] = strings.Split(line, ",")
	}

	return rules, updates
}
