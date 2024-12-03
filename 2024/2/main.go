package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var SAFE_DIFF int = 3

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
	lines := strings.Split(input, "\n")

	var safeReports int

	for i, line := range lines {
		report := parseReport((line))
		isReportSafe := checkReport(i, report, 0)

		if isReportSafe {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)
}

func parseReport(line string) []int {
	chars := strings.Split(line, " ")
	report := make([]int, len(chars))

	for i, char := range strings.Split(line, " ") {
		reading, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println(err.Error())
			panic("Incorrect input")
		}

		report[i] = reading
	}
	return report
}

func checkReport(reportNum int, report []int, allowedFails int) bool {
	slog.Debug(fmt.Sprintf("Checking report #%d: %v. Allowed to fail %d times", reportNum, report, allowedFails))

	direction := report[1] - report[0]
	if direction < 0 {
		slog.Debug("Direction: descending")
	}
	if direction > 0 {
		slog.Debug("Direction: ascending")
	}
	for i := 0; i < (len(report) - 1); i++ {
		diff := float64(report[i+1] - report[i])
		isDiffSafe := math.Abs(diff) <= float64(SAFE_DIFF) && math.Abs(diff) > 0

		if (!isDiffSafe || direction > 0 && diff < 0) || (direction < 0 && diff > 0) {
			slog.Debug(fmt.Sprintf("Problematic levels: %v, %v", report[i], report[i+1]))
			if allowedFails > 0 {
				slog.Debug("Allowed to fail")
				for i := 0; i < len(report); i++ {
					modifiedReport := slices.Delete(append([]int(nil), report...), i, i+1)
					slog.Debug(fmt.Sprintf("About to check report: %v", modifiedReport))
					result := checkReport(reportNum, modifiedReport, allowedFails-1)
					if result {
						slog.Debug(fmt.Sprintf("Report #%d is safe. No further checks required", reportNum))
						return true
					} else {
						slog.Debug("Modified report is not safe")
						continue
					}
				}
				slog.Debug("None of the modified reports is safe. Report is unsafe")
				return false
			}
			slog.Debug("Not allowed to fail")
			slog.Debug(fmt.Sprintf("Report #%d is not safe", reportNum))
			return false
		}
	}
	slog.Debug(fmt.Sprintf("Report #%d is safe", reportNum))

	return true

}

func secondPart() {
	lines := strings.Split(input, "\n")

	var safeReports int

	for i, line := range lines {
		report := parseReport((line))
		isReportSafe := checkReport(i, report, 1)

		if isReportSafe {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)
}
