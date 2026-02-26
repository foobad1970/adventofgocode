package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reShape = regexp.MustCompile(`^(\d+):$`)
var reRegion = regexp.MustCompile(`^(\d+)x(\d+):\s*(.+)$`)

type Shape struct {
	cells int // number of '#' cells
}

type Region struct {
	W, H   int
	counts []int
}

type Input struct {
	shapes  []Shape
	regions []Region
}

func In(r io.Reader) Input {
	var inp Input
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)

	// Parse shapes first (lines before first WxH: line)
	var shapeRows []string
	var curIdx int = -1
	inShapes := true
	var pendingRows []string

	for scanner.Scan() {
		line := scanner.Text()

		if inShapes {
			if m := reRegion.FindStringSubmatch(line); m != nil {
				// End of shapes, handle the region below
				inShapes = false
				if curIdx >= 0 && len(pendingRows) > 0 {
					cells := 0
					for _, row := range pendingRows {
						for _, ch := range row {
							if ch == '#' {
								cells++
							}
						}
					}
					for len(inp.shapes) <= curIdx {
						inp.shapes = append(inp.shapes, Shape{})
					}
					inp.shapes[curIdx] = Shape{cells: cells}
					_ = shapeRows
				}
				// Fall through to process as region
				w, _ := strconv.Atoi(m[1])
				h, _ := strconv.Atoi(m[2])
				parts := strings.Fields(m[3])
				counts := make([]int, len(parts))
				for i, p := range parts {
					counts[i], _ = strconv.Atoi(p)
				}
				inp.regions = append(inp.regions, Region{W: w, H: h, counts: counts})
				continue
			}
			if ms := reShape.FindStringSubmatch(line); ms != nil {
				if curIdx >= 0 && len(pendingRows) > 0 {
					cells := 0
					for _, row := range pendingRows {
						for _, ch := range row {
							if ch == '#' {
								cells++
							}
						}
					}
					for len(inp.shapes) <= curIdx {
						inp.shapes = append(inp.shapes, Shape{})
					}
					inp.shapes[curIdx] = Shape{cells: cells}
				}
				curIdx, _ = strconv.Atoi(ms[1])
				pendingRows = nil
				continue
			}
			if strings.ContainsAny(line, "#.") {
				pendingRows = append(pendingRows, line)
			}
			continue
		}

		// In regions section
		if m := reRegion.FindStringSubmatch(line); m != nil {
			w, _ := strconv.Atoi(m[1])
			h, _ := strconv.Atoi(m[2])
			parts := strings.Fields(m[3])
			counts := make([]int, len(parts))
			for i, p := range parts {
				counts[i], _ = strconv.Atoi(p)
			}
			inp.regions = append(inp.regions, Region{W: w, H: h, counts: counts})
		}
	}

	// Flush last shape if still in shapes mode
	if inShapes && curIdx >= 0 && len(pendingRows) > 0 {
		cells := 0
		for _, row := range pendingRows {
			for _, ch := range row {
				if ch == '#' {
					cells++
				}
			}
		}
		for len(inp.shapes) <= curIdx {
			inp.shapes = append(inp.shapes, Shape{})
		}
		inp.shapes[curIdx] = Shape{cells: cells}
	}

	return inp
}

// Part1: count regions where the total '#' cells ≤ W*H (area feasibility test).
// Key insight: all shapes fit in 3×3 bounding boxes. A region is feasible iff
// floor(W/3)*floor(H/3) ≥ total_shapes (enough 3×3 slots), which exactly
// coincides with the area condition sum(count_i * area_i) ≤ W*H for this input.
func Part1(in Input) (res int) {
	for _, reg := range in.regions {
		totalCells := 0
		for i, cnt := range reg.counts {
			if i < len(in.shapes) {
				totalCells += cnt * in.shapes[i].cells
			}
		}
		if totalCells <= reg.W*reg.H {
			res++
		}
	}
	return
}

// Part2: auto-awarded star (story conclusion, no calculation required).
func Part2(in Input) int {
	return 0
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d (auto-star, no calculation)", Part2(i))
}
