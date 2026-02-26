package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024*4), 1024*1024*4)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return
}

type problem struct {
	op   byte
	nums []int64
}

func parseProblems(in Input) []problem {
	if len(in) == 0 {
		return nil
	}
	width := 0
	for _, line := range in {
		if len(line) > width {
			width = len(line)
		}
	}
	// pad all lines to same width
	padded := make([]string, len(in))
	for i, line := range in {
		if len(line) < width {
			padded[i] = line + strings.Repeat(" ", width-len(line))
		} else {
			padded[i] = line
		}
	}

	opRow := padded[len(padded)-1]
	numRows := padded[:len(padded)-1]

	// find all-space columns
	allSpace := make([]bool, width)
	for c := 0; c < width; c++ {
		allSpace[c] = true
		for _, row := range padded {
			if row[c] != ' ' {
				allSpace[c] = false
				break
			}
		}
	}

	// find column groups
	var problems []problem
	i := 0
	for i < width {
		if !allSpace[i] {
			j := i
			for j < width && !allSpace[j] {
				j++
			}
			// group is [i, j-1]
			op := byte('+')
			seg := strings.TrimSpace(opRow[i:j])
			if len(seg) > 0 {
				op = seg[0]
			}
			var nums []int64
			for _, row := range numRows {
				s := strings.TrimSpace(row[i:j])
				if s != "" {
					n, _ := strconv.ParseInt(s, 10, 64)
					nums = append(nums, n)
				}
			}
			problems = append(problems, problem{op, nums})
			i = j
		} else {
			i++
		}
	}
	return problems
}

func solve(p problem) int64 {
	if len(p.nums) == 0 {
		return 0
	}
	res := p.nums[0]
	for _, n := range p.nums[1:] {
		if p.op == '+' {
			res += n
		} else {
			res *= n
		}
	}
	return res
}

func Part1(in Input) (res int) {
	problems := parseProblems(in)
	var total int64
	for _, p := range problems {
		total += solve(p)
	}
	return int(total)
}

func Part2(in Input) (res int) {
	if len(in) == 0 {
		return
	}
	width := 0
	for _, line := range in {
		if len(line) > width {
			width = len(line)
		}
	}
	padded := make([]string, len(in))
	for i, line := range in {
		if len(line) < width {
			padded[i] = line + strings.Repeat(" ", width-len(line))
		} else {
			padded[i] = line
		}
	}

	opRow := padded[len(padded)-1]
	numRows := padded[:len(padded)-1]

	allSpace := make([]bool, width)
	for c := 0; c < width; c++ {
		allSpace[c] = true
		for _, row := range padded {
			if row[c] != ' ' {
				allSpace[c] = false
				break
			}
		}
	}

	var total int64
	i := 0
	for i < width {
		if !allSpace[i] {
			j := i
			for j < width && !allSpace[j] {
				j++
			}
			// group columns [i, j-1], right-to-left
			seg := strings.TrimSpace(opRow[i:j])
			op := byte('+')
			if len(seg) > 0 {
				op = seg[0]
			}

			var nums []int64
			for c := j - 1; c >= i; c-- {
				// collect digits top-to-bottom
				var digits []byte
				for _, row := range numRows {
					ch := row[c]
					if ch >= '0' && ch <= '9' {
						digits = append(digits, ch)
					}
				}
				if len(digits) > 0 {
					n, _ := strconv.ParseInt(string(digits), 10, 64)
					nums = append(nums, n)
				}
			}

			result := solve(problem{op, nums})
			total += result
			i = j
		} else {
			i++
		}
	}
	return int(total)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
