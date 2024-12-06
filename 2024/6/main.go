package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"
)

type PatrolMap [][]string

func (pm PatrolMap) String() string {
	var b strings.Builder
	for _, line := range pm {
		b.WriteString(fmt.Sprintf("%v\n", line))
	}
	return b.String()
}

func (pm PatrolMap) getVisitedPoints() []Position {
	var visited []Position
	visitedSymbols := append(slices.Collect(maps.Values(directionSymbols)), "X")
	for i, line := range pm {
		for j, point := range line {
			if slices.Contains(visitedSymbols, point) {
				visited = append(visited, Position{i, j})
			}
		}
	}
	return visited
}

type Position [2]int

type Direction int

//go:embed input.txt
var input string

const (
	North Direction = iota
	East
	South
	West
)

var directionSymbols = map[Direction]string{
	North: "^",
	East:  ">",
	South: "V",
	West:  "<",
}

var increments = map[Direction][2]int{
	East:  {0, 1},
	West:  {0, -1},
	North: {-1, 0},
	South: {1, 0},
}

func (d Direction) String() string {
	return directionSymbols[d]
}

func (d Direction) GetIncrement() [2]int {
	return increments[d]
}

type Guard struct {
	position  Position
	direction Direction
	patrolMap PatrolMap
	isActive  bool
}

func newGuard(p Position, pm PatrolMap) *Guard {

	return &Guard{position: p, direction: North, patrolMap: pm, isActive: true}
}

func (g *Guard) move() {
	inc := (*g).direction.GetIncrement()
	(*g).patrolMap[(*g).position[0]][(*g).position[1]] = "X"
	(*g).position[0], (*g).position[1] = (*g).position[0]+inc[0], (*g).position[1]+inc[1]
	(*g).patrolMap[(*g).position[0]][(*g).position[1]] = (*g).direction.String()
}

func (g *Guard) rotate() {
	switch (*g).direction {
	case North:
		(*g).direction = East
	case East:
		(*g).direction = South
	case South:
		(*g).direction = West
	case West:
		(*g).direction = North
	}
}

func (g *Guard) checkSpaceAhead() bool {
	inc := (*g).direction.GetIncrement()
	positionToCheck := Position{(*g).position[0] + inc[0], (*g).position[1] + inc[1]}

	if positionToCheck[0] == len((*g).patrolMap) || positionToCheck[1] == len((*g).patrolMap[0]) || positionToCheck[0] == -1 || positionToCheck[1] == -1 {
		(*g).isActive = false
		return false
	}

	return (*g).patrolMap[positionToCheck[0]][positionToCheck[1]] != "#"
}

func (g *Guard) patrol() {
	for (*g).isActive {
		if (*g).checkSpaceAhead() && g.isActive {
			(*g).move()
			continue
		}
		(*g).rotate()
	}
}

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

	pm, sp := prepareInput(input)
	guard := *newGuard(sp, pm)

	guard.patrol()
	fmt.Println(len(guard.patrolMap.getVisitedPoints()))
}

func secondPart() {
	slog.Debug("Running second part")

}

func prepareInput(input string) (PatrolMap, Position) {
	lines := strings.Split(input, "\n")
	preparedInput := make([][]string, len(lines))
	var startingPosition Position
	for i, line := range lines {
		preparedInput[i] = strings.Split(line, "")
		for j, point := range preparedInput[i] {
			if point == "^" {
				startingPosition = Position{i, j}
			}
		}
	}

	return preparedInput, startingPosition
}
