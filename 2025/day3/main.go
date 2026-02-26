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

func maxJoltage2(bank string) int {
	best := 0
	n := len(bank)
	maxRight := make([]int, n)
	maxRight[n-1] = int(bank[n-1] - '0')
	for i := n - 2; i >= 0; i-- {
		d := int(bank[i] - '0')
		if d > maxRight[i+1] {
			maxRight[i] = d
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}
	for i := 0; i < n-1; i++ {
		d := int(bank[i] - '0')
		val := d*10 + maxRight[i+1]
		if val > best {
			best = val
		}
	}
	return best
}

// selectK picks k digits from bank (in order) to form the largest number.
func selectK(bank string, k int) int {
	n := len(bank)
	result := 0
	start := 0
	for i := 0; i < k; i++ {
		end := n - k + i
		best := start
		for j := start + 1; j <= end; j++ {
			if bank[j] > bank[best] {
				best = j
			}
		}
		result = result*10 + int(bank[best]-'0')
		start = best + 1
	}
	return result
}

func Part1(in Input) (res int) {
	for _, bank := range in {
		res += maxJoltage2(bank)
	}
	return
}

func Part2(in Input) (res int) {
	for _, bank := range in {
		res += selectK(bank, 12)
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
