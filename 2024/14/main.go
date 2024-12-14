package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func isPosEqual(a Position, b Position) bool {
	return a.x == b.x && a.y == b.y
}

func newPosition(s string) Position {
	var posInts []int
	posStrs := strings.Split(s, ",")
	for _, ps := range posStrs {
		num, _ := strconv.Atoi(ps)
		posInts = append(posInts, num)
	}

	return Position{x: posInts[0], y: posInts[1]}
}

type Robot struct {
	Position Position
	velocity struct {
		vx int
		vy int
	}
}

func (r *Robot) Move(times int) {
	newX := (r.Position.x + r.velocity.vx*times) % bounds[0]
	if newX < 0 {
		newX = bounds[0] + newX
	}
	newY := (r.Position.y + r.velocity.vy*times) % bounds[1]
	if newY < 0 {
		newY = bounds[1] + newY
	}
	r.Position = Position{x: newX, y: newY}
}

func newRobot(s string) *Robot {
	parts := strings.Fields(s)
	for i, p := range parts {
		parts[i] = strings.Split(p, "=")[1]
	}

	var velInts []int
	velStrs := strings.Split(parts[1], ",")
	for _, vs := range velStrs {
		num, _ := strconv.Atoi(vs)
		velInts = append(velInts, num)
	}

	return &Robot{
		Position: newPosition(parts[0]),
		velocity: struct {
			vx int
			vy int
		}{
			vx: velInts[0],
			vy: velInts[1],
		},
	}
}

//go:embed input.txt
var input string
var bounds = [2]int{101, 103}

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

	strs := strings.Split(input, "\n")
	var robots []*Robot
	for _, s := range strs {
		robots = append(robots, newRobot(s))
	}

	var q1, q2, q3, q4 []*Robot
	for _, r := range robots {
		r.Move(100)

		switch {
		case r.Position.x < (bounds[0]/2) && r.Position.y < (bounds[1]/2):
			q1 = append(q1, r)
		case r.Position.x > (bounds[0]/2) && r.Position.y < (bounds[1]/2):
			q2 = append(q2, r)
		case r.Position.x < (bounds[0]/2) && r.Position.y > (bounds[1]/2):
			q3 = append(q3, r)
		case r.Position.x > (bounds[0]/2) && r.Position.y > (bounds[1]/2):
			q4 = append(q4, r)
		default:
			continue
		}
	}
	safetyFactor := len(q1) * len(q2) * len(q3) * len(q4)

	fmt.Println(safetyFactor)
}

func secondPart() {
	slog.Debug("Running second part")
}
