package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			res = append(res, line)
		}
	}

	return
}

func rotate(pos, amount int, dir byte) int {
	if dir == 'L' {
		pos = ((pos - amount) % 100 + 100) % 100
	} else {
		pos = (pos + amount) % 100
	}
	return pos
}

func Part1(in Input) (res int) {
	pos := 50
	for _, line := range in {
		dir := line[0]
		var amount int
		fmt.Sscanf(line[1:], "%d", &amount)
		pos = rotate(pos, amount, dir)
		if pos == 0 {
			res++
		}
	}
	return
}

func countZeros(pos, amount int, dir byte) int {
	var offset int
	if dir == 'R' {
		offset = (100 - pos) % 100
		if offset == 0 {
			offset = 100
		}
	} else {
		offset = pos
		if offset == 0 {
			offset = 100
		}
	}
	if amount < offset {
		return 0
	}
	return (amount-offset)/100 + 1
}

func Part2(in Input) (res int) {
	pos := 50
	for _, line := range in {
		dir := line[0]
		var amount int
		fmt.Sscanf(line[1:], "%d", &amount)
		res += countZeros(pos, amount, dir)
		pos = rotate(pos, amount, dir)
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
