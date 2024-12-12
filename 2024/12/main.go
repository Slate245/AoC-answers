package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

type PlotNode struct {
	x          int
	y          int
	plantType  string
	neighbours []*PlotNode
}

func (p *PlotNode) GetNeighboursOfSameType() []*PlotNode {
	var neighboursOfSameType []*PlotNode
	for _, neighbour := range p.neighbours {
		if neighbour.plantType != p.plantType {
			continue
		}
		neighboursOfSameType = append(neighboursOfSameType, neighbour)
	}
	return neighboursOfSameType
}

type Region struct {
	plantType string
	plots     []*PlotNode
	area      int
	perimeter int
}

func (r *Region) AddPlot(p *PlotNode) {
	if p.plantType != r.plantType {
		panic(fmt.Sprintf("can't add plot of type '%s' to region of type '%s'", p.plantType, r.plantType))
	}
	neighboursOfSameType := p.GetNeighboursOfSameType()

	r.plots = append(r.plots, neighboursOfSameType...)
	r.area += 1
	r.perimeter += 4 - len(neighboursOfSameType)
}

func (r *Region) GetFencePrice() int {
	return r.area * r.perimeter
}

func newRegion(p *PlotNode) *Region {
	if p == nil {
		panic("can't construct region for nil")
	}

	plantType := p.plantType
	var plots []*PlotNode
	area := 0
	perimeter := 0
	r := Region{plantType, plots, area, perimeter}
	queue := []*PlotNode{p}
	var visited []*PlotNode

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if slices.Contains(visited, node) {
			continue
		}

		visited = append(visited, node)
		r.AddPlot(node)

		queue = append(queue, node.GetNeighboursOfSameType()...)
	}

	return &r
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

	nodes := prepareInput(input)
	var visitedNodes []*PlotNode
	var regions []*Region

	for _, node := range nodes {
		if slices.Contains(visitedNodes, node) {
			continue
		}

		r := newRegion(node)
		visitedNodes = append(visitedNodes, r.plots...)

		regions = append(regions, newRegion(node))
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.GetFencePrice()
	}

	fmt.Println(totalPrice)
}

func secondPart() {
	slog.Debug("Running second part")
}

func prepareInput(input string) []*PlotNode {
	var preparedInput []*PlotNode
	var plotMap = make(map[string]*PlotNode)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		symbols := strings.Split(line, "")
		for x, plantType := range symbols {
			var neighbours []*PlotNode
			pn := PlotNode{x, y, plantType, neighbours}
			plotMap[fmt.Sprint([]int{x, y})] = &pn
			preparedInput = append(preparedInput, &pn)
		}
	}

	for _, plotNode := range preparedInput {
		x := plotNode.x
		y := plotNode.y
		neighbourCoords := [][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}

		for _, coords := range neighbourCoords {
			key := fmt.Sprint(coords)

			neighbour, neighbourExists := plotMap[key]
			if !neighbourExists {
				continue
			}
			plotNode.neighbours = append(plotNode.neighbours, neighbour)
		}
	}

	return preparedInput
}
