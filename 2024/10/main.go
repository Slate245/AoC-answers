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
	x      int
	y      int
	height int
	left   *Position
	right  *Position
	top    *Position
	bottom *Position
}

func (p *Position) GetScore() int {
	if p.height != 0 {
		return 0
	}
	positionsToCheck := make(map[string]*Position)
	positionsToCheck[p.StrCoords()] = p

	for i := 1; i < 10; i++ {
		nextPositions := make(map[string]*Position)
		for _, pos := range positionsToCheck {
			if pos.top != nil && pos.top.height == i {
				nextPositions[pos.top.StrCoords()] = pos.top
			}
			if pos.right != nil && pos.right.height == i {
				nextPositions[pos.right.StrCoords()] = pos.right
			}
			if pos.bottom != nil && pos.bottom.height == i {
				nextPositions[pos.bottom.StrCoords()] = pos.bottom
			}
			if pos.left != nil && pos.left.height == i {
				nextPositions[pos.left.StrCoords()] = pos.left
			}
		}
		positionsToCheck = nextPositions
	}
	return len(positionsToCheck)
}

func (p *Position) GetRating() int {
	if p.height != 0 {
		return 0
	}

	paths := [][]*Position{{p}}

	for i := 1; i < 10; i++ {
		var nextPaths [][]*Position
		for _, p := range paths {
			if len(p) == 0 {
				continue
			}
			pos := p[len(p)-1]
			if pos.top != nil && pos.top.height == i {
				var newPath []*Position
				copy(newPath, p)
				newPath = append(newPath, pos.top)
				nextPaths = append(nextPaths, newPath)
			}
			if pos.right != nil && pos.right.height == i {
				var newPath []*Position
				copy(newPath, p)
				newPath = append(newPath, pos.right)
				nextPaths = append(nextPaths, newPath)
			}
			if pos.bottom != nil && pos.bottom.height == i {
				var newPath []*Position
				copy(newPath, p)
				newPath = append(newPath, pos.bottom)
				nextPaths = append(nextPaths, newPath)
			}
			if pos.left != nil && pos.left.height == i {
				var newPath []*Position
				copy(newPath, p)
				newPath = append(newPath, pos.left)
				nextPaths = append(nextPaths, newPath)
			}
		}
		paths = nextPaths
	}
	return len(paths)
}

func (p *Position) StrCoords() string {
	return fmt.Sprintf("{%d:%d}", p.x, p.y)
}

func (p *Position) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d{", p.height))
	if p.top != nil {
		b.WriteString("t:")
		b.WriteString(p.top.StrCoords())
	}
	if p.right != nil {
		b.WriteString("r:")
		b.WriteString(p.right.StrCoords())
	}
	if p.bottom != nil {
		b.WriteString("b:")
		b.WriteString(p.bottom.StrCoords())
	}
	if p.left != nil {
		b.WriteString("l:")
		b.WriteString(p.left.StrCoords())
	}
	b.WriteRune('}')
	return b.String()
}

type TopographicMap struct {
	tmap   map[string]*Position
	bounds [2][2]int
}

func (tm *TopographicMap) set(x int, y int, p *Position) {
	tm.updateBounds(p)
	p.left = tm.get(x-1, y)
	p.right = tm.get(x+1, y)
	p.top = tm.get(x, y-1)
	p.bottom = tm.get(x, y+1)
	if p.top != nil {
		p.top.bottom = p
	}
	if p.right != nil {
		p.right.left = p
	}
	if p.left != nil {
		p.left.right = p
	}
	if p.bottom != nil {
		p.bottom.top = p
	}
	key := fmt.Sprintf("%d:%d", x, y)
	tm.tmap[key] = p
}

func (tm *TopographicMap) updateBounds(p *Position) {
	if p == nil {
		return
	}
	if p.x < tm.bounds[0][0] {
		tm.bounds[0][0] = p.x
	}
	if p.x > tm.bounds[0][1] {
		tm.bounds[0][1] = p.x
	}
	if p.y < tm.bounds[1][0] {
		tm.bounds[1][0] = p.y
	}
	if p.y > tm.bounds[1][1] {
		tm.bounds[1][1] = p.y
	}
}

func (tm *TopographicMap) get(x int, y int) *Position {
	key := fmt.Sprintf("%d:%d", x, y)
	p, ok := tm.tmap[key]

	if !ok {
		return nil
	}
	return p
}

func (tm *TopographicMap) String() string {
	var b strings.Builder
	for y := tm.bounds[1][0]; y < tm.bounds[1][1]+1; y++ {
		for x := tm.bounds[0][0]; x < tm.bounds[0][1]+1; x++ {
			p := tm.get(x, y)

			if p == nil {
				b.WriteRune('.')
				continue
			}
			b.WriteString(fmt.Sprint(p))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func newTopographicMap(input string) *TopographicMap {
	tmap := make(map[string]*Position)
	bounds := [2][2]int{{0, 0}, {0, 0}}
	tm := TopographicMap{tmap, bounds}

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		symbols := strings.Split(line, "")
		for x, symbol := range symbols {
			height, err := strconv.Atoi(symbol)

			if err != nil {
				height = -1
			}
			var left *Position = nil
			var right *Position = nil
			var top *Position = nil
			var bottom *Position = nil

			tm.set(x, y, &Position{x, y, height, left, right, top, bottom})
		}
	}
	return &tm
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

	tmap := prepareInput(input)
	totalScore := 0

	for y := tmap.bounds[1][0]; y < tmap.bounds[1][1]+1; y++ {
		for x := tmap.bounds[0][0]; x < tmap.bounds[0][1]+1; x++ {
			p := tmap.get(x, y)

			if p == nil {
				continue
			}
			totalScore += p.GetScore()
		}
	}

	fmt.Println(totalScore)
}

func secondPart() {
	slog.Debug("Running second part")

	tmap := prepareInput(input)
	totalRating := 0

	for y := tmap.bounds[1][0]; y < tmap.bounds[1][1]+1; y++ {
		for x := tmap.bounds[0][0]; x < tmap.bounds[0][1]+1; x++ {
			p := tmap.get(x, y)

			if p == nil {
				continue
			}
			totalRating += p.GetRating()
		}
	}

	fmt.Println(totalRating)
}

func prepareInput(input string) TopographicMap {
	return *newTopographicMap(input)
}
