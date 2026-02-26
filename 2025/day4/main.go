package main

import (
	"bufio"
	"io"
	"log"
	"os"
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

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func neighbors(grid Input, r, c int) int {
	count := 0
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[nr]) {
			if grid[nr][nc] == '@' {
				count++
			}
		}
	}
	return count
}

func Part1(in Input) (res int) {
	for r, row := range in {
		for c := range row {
			if row[c] == '@' && neighbors(in, r, c) < 4 {
				res++
			}
		}
	}
	return
}

func Part2(in Input) (res int) {
	rows := len(in)
	cols := len(in[0])
	// copy grid into mutable form
	grid := make([][]byte, rows)
	for r, row := range in {
		grid[r] = []byte(row)
	}

	countNeighbors := func(r, c int) int {
		count := 0
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
				count++
			}
		}
		return count
	}

	type pos struct{ r, c int }
	queue := []pos{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' && countNeighbors(r, c) < 4 {
				queue = append(queue, pos{r, c})
			}
		}
	}

	inQueue := make([][]bool, rows)
	for r := range inQueue {
		inQueue[r] = make([]bool, cols)
	}
	for _, p := range queue {
		inQueue[p.r][p.c] = true
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if grid[p.r][p.c] != '@' {
			continue
		}
		grid[p.r][p.c] = '.'
		res++
		// check neighbors
		for _, d := range dirs {
			nr, nc := p.r+d[0], p.c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
				grid[nr][nc] == '@' && !inQueue[nr][nc] &&
				countNeighbors(nr, nc) < 4 {
				inQueue[nr][nc] = true
				queue = append(queue, pos{nr, nc})
			}
		}
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
