package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"strings"
)

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

	freqToTowersMap, bounds := prepareInput(input)

	result := make(map[string]bool)

	for _, towers := range freqToTowersMap {
		towerIdxPairs := getIndexPairs(len(towers))
		for _, idxs := range towerIdxPairs {
			antinodes := getAntinodes(towers[idxs[0]], towers[idxs[1]], bounds)

			for _, an := range antinodes {
				key := fmt.Sprint(an)
				result[key] = true
			}
		}
	}

	fmt.Println(len(result))

}

func secondPart() {
	slog.Debug("Running second part")
}

func prepareInput(input string) (map[string][][2]int, [2]int) {
	lines := strings.Split(input, "\n")
	freqToTowerCoordsMap := make(map[string][][2]int)
	bounds := [2]int{len(lines[0]), len(lines)}
	for yCoord, line := range lines {
		symbols := strings.Split(line, "")
		for xCoord, symbol := range symbols {
			if symbol == "." {
				continue
			}
			coords := freqToTowerCoordsMap[symbol]
			coords = append(coords, [2]int{xCoord, yCoord})
			freqToTowerCoordsMap[symbol] = coords
		}
	}

	return freqToTowerCoordsMap, bounds
}

func getIndexPairs(len int) [][2]int {
	var result [][2]int
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			result = append(result, [2]int{i, j})
		}
	}
	return result
}

func getAntinodes(t1 [2]int, t2 [2]int, bounds [2]int) [][2]int {
	delta := [2]int{t1[0] - t2[0], t1[1] - t2[1]}
	var antinodes [][2]int
	an1 := [2]int{t1[0] + delta[0], t1[1] + delta[1]}
	an2 := [2]int{t2[0] - delta[0], t2[1] - delta[1]}

	if an1[0] < bounds[0] && an1[0] >= 0 && an1[1] < bounds[1] && an1[1] >= 0 {
		antinodes = append(antinodes, an1)
	}

	if an2[0] < bounds[0] && an2[0] >= 0 && an2[1] < bounds[1] && an2[1] >= 0 {
		antinodes = append(antinodes, an2)
	}

	return antinodes
}
