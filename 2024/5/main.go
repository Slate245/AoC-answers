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
	validUpdates, _ := checkUpdatesWithRules(updates, rules)

	result := getResultFromUpdates(validUpdates)
	fmt.Println(result)
}

func secondPart() {
	slog.Debug("Running second part")
	rules, updates := prepareInput(input)
	_, invalidUpdates := checkUpdatesWithRules(updates, rules)
	slog.Debug(fmt.Sprintf("Invalid updates: %v", invalidUpdates))

	for _, update := range invalidUpdates {
		update = *fixUpdate(&update, rules)
	}
	result := getResultFromUpdates(invalidUpdates)
	fmt.Println(result)
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

func checkUpdatesWithRules(updates [][]string, rules [][2]string) ([][]string, [][]string) {
	var validUpdates [][]string
	var invalidUpdates [][]string

	for _, update := range updates {
		if isValid, _ := validateUpdate(update, rules); isValid {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	slog.Debug(fmt.Sprintf("Valid updates: %v", validUpdates))
	slog.Debug(fmt.Sprintf("Found %d valid updates", len(validUpdates)))
	return validUpdates, invalidUpdates
}

func validateUpdate(update []string, rules [][2]string) (bool, [2]int) {
	slog.Debug(fmt.Sprintf("Checking update: %v", update))
	isValid := true
	var errIndices [2]int
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
			errIndices = [2]int{firstIndex, secondIndex}
			break
		}
	}
	return isValid, errIndices
}

func getResultFromUpdates(validUpdates [][]string) int {
	result := 0
	for _, update := range validUpdates {
		middleNumber, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		slog.Debug(fmt.Sprintf("Middle number in %v is %d", update, middleNumber))
		result += middleNumber
	}
	return result
}

func fixUpdate(update *[]string, rules [][2]string) *[]string {
	if isValid, errIndices := validateUpdate(*update, rules); !isValid {
		(*update)[errIndices[0]], (*update)[errIndices[1]] = (*update)[errIndices[1]], (*update)[errIndices[0]]
		fixUpdate(update, rules)
	}
	return update
}
