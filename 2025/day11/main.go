package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Input map[string][]string

func In(r io.Reader) Input {
	graph := make(Input)
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		idx := strings.Index(line, ": ")
		if idx < 0 {
			continue
		}
		node := line[:idx]
		children := strings.Fields(line[idx+2:])
		graph[node] = children
	}
	return graph
}

// Part 1: count all paths from "you" to "out"
func Part1(in Input) int {
	memo := make(map[string]int)
	var count func(node string) int
	count = func(node string) int {
		if node == "out" {
			return 1
		}
		if v, ok := memo[node]; ok {
			return v
		}
		total := 0
		for _, c := range in[node] {
			total += count(c)
		}
		memo[node] = total
		return total
	}
	return count("you")
}

// Part 2: count paths from "svr" to "out" that visit both "dac" and "fft"
// mask: bit 0 = dac visited, bit 1 = fft visited
func Part2(in Input) int {
	type state struct {
		node string
		mask int
	}
	memo := make(map[state]int)
	var count func(node string, mask int) int
	count = func(node string, mask int) int {
		if node == "out" {
			if mask == 3 {
				return 1
			}
			return 0
		}
		s := state{node, mask}
		if v, ok := memo[s]; ok {
			return v
		}
		total := 0
		for _, c := range in[node] {
			nm := mask
			if c == "dac" {
				nm |= 1
			}
			if c == "fft" {
				nm |= 2
			}
			total += count(c, nm)
		}
		memo[s] = total
		return total
	}
	return count("svr", 0)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
