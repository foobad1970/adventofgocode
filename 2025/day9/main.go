package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Point struct{ x, y int64 }

type Input []Point

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var p Point
		fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		res = append(res, p)
	}
	return
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func Part1(in Input) (res int) {
	n := len(in)
	var best int64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area := (abs(in[i].x-in[j].x) + 1) * (abs(in[i].y-in[j].y) + 1)
			if area > best {
				best = area
			}
		}
	}
	return int(best)
}

func sortedUnique(vals []int64) []int64 {
	seen := make(map[int64]bool)
	var u []int64
	for _, v := range vals {
		if !seen[v] {
			seen[v] = true
			u = append(u, v)
		}
	}
	sort.Slice(u, func(i, j int) bool { return u[i] < u[j] })
	return u
}

func Part2(in Input) (res int) {
	n := len(in)

	// build horizontal edges of the polygon
	type hedge struct{ y, xlo, xhi int64 }
	var hedges []hedge
	for i := 0; i < n; i++ {
		a, b := in[i], in[(i+1)%n]
		if a.y == b.y {
			xlo, xhi := a.x, b.x
			if xlo > xhi {
				xlo, xhi = xhi, xlo
			}
			hedges = append(hedges, hedge{a.y, xlo, xhi})
		}
	}

	// coordinate compression
	ux := sortedUnique(func() []int64 {
		v := make([]int64, n)
		for i, p := range in {
			v[i] = p.x
		}
		return v
	}())
	uy := sortedUnique(func() []int64 {
		v := make([]int64, n)
		for i, p := range in {
			v[i] = p.y
		}
		return v
	}())
	p, q := len(ux), len(uy)

	uxIdx := make(map[int64]int, p)
	for i, x := range ux {
		uxIdx[x] = i
	}
	uyIdx := make(map[int64]int, q)
	for j, y := range uy {
		uyIdx[y] = j
	}

	// for gap cell (i,j): test point tx2 = ux[i]+ux[i+1], ty2 = uy[j]+uy[j+1]
	// count h-edges above ty2 crossing tx2 (strictly)
	isInsideGap := func(i, j int) bool {
		tx2 := ux[i] + ux[i+1]
		ty2 := uy[j] + uy[j+1]
		count := 0
		for _, e := range hedges {
			if 2*e.y > ty2 && 2*e.xlo < tx2 && tx2 < 2*e.xhi {
				count++
			}
		}
		return count%2 == 1
	}

	// precompute inside_gap[i][j]
	insideGap := make([][]bool, p-1)
	for i := range insideGap {
		insideGap[i] = make([]bool, q-1)
		for j := range insideGap[i] {
			insideGap[i][j] = isInsideGap(i, j)
		}
	}

	// 2D prefix sum (p×q, 0-indexed, prefix[i+1][j+1] covers insideGap[0..i][0..j])
	prefix := make([][]int, p)
	for i := range prefix {
		prefix[i] = make([]int, q)
	}
	for i := 0; i < p-1; i++ {
		for j := 0; j < q-1; j++ {
			v := 0
			if insideGap[i][j] {
				v = 1
			}
			prefix[i+1][j+1] = v + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}

	queryGaps := func(ix1, ix2, iy1, iy2 int) (inside, total int) {
		if ix1 >= ix2 || iy1 >= iy2 {
			return 0, 0
		}
		total = (ix2 - ix1) * (iy2 - iy1)
		inside = prefix[ix2][iy2] - prefix[ix1][iy2] - prefix[ix2][iy1] + prefix[ix1][iy1]
		return
	}

	var best int64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1, y1 := in[i].x, in[i].y
			x2, y2 := in[j].x, in[j].y
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			if x1 == x2 || y1 == y2 {
				continue
			}
			ix1, ix2 := uxIdx[x1], uxIdx[x2]
			iy1, iy2 := uyIdx[y1], uyIdx[y2]
			inside, total := queryGaps(ix1, ix2, iy1, iy2)
			if inside != total {
				continue
			}
			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area > best {
				best = area
			}
		}
	}
	return int(best)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
