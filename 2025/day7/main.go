package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			res = append(res, line)
		}
	}
	return
}

type pos struct{ c, r int }

func simulate(grid Input, startC, startR int) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[pos]bool)
	queue := []pos{{startC, startR}}
	splits := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.r < 0 || p.r >= rows || p.c < 0 || p.c >= cols {
			continue
		}
		if visited[p] {
			continue
		}
		visited[p] = true

		ch := grid[p.r][p.c]
		if ch == '^' {
			splits++
			queue = append(queue, pos{p.c - 1, p.r}, pos{p.c + 1, p.r})
		} else {
			queue = append(queue, pos{p.c, p.r + 1})
		}
	}
	return splits
}

func Part1(in Input) (res int) {
	for r, row := range in {
		c := strings.IndexByte(row, 'S')
		if c >= 0 {
			res = simulate(in, c, r)
			return
		}
	}
	return
}

func countTimelines(grid Input, c, r int, memo map[pos]int) int {
	rows := len(grid)
	cols := len(grid[0])
	if c < 0 || c >= cols {
		return 0
	}
	if r >= rows {
		return 1
	}
	p := pos{c, r}
	if v, ok := memo[p]; ok {
		return v
	}
	var result int
	if grid[r][c] == '^' {
		result = countTimelines(grid, c-1, r, memo) + countTimelines(grid, c+1, r, memo)
	} else {
		result = countTimelines(grid, c, r+1, memo)
	}
	memo[p] = result
	return result
}

func Part2(in Input) (res int) {
	memo := make(map[pos]int)
	for r, row := range in {
		c := strings.IndexByte(row, 'S')
		if c >= 0 {
			return countTimelines(in, c, r, memo)
		}
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
