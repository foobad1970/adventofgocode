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
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			res = append(res, line)
		}
	}

	return
}

type rng struct{ lo, hi int64 }

func parseRanges(in Input) []rng {
	var ranges []rng
	for _, line := range in {
		for _, part := range strings.Split(strings.TrimSpace(line), ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			dash := strings.Index(part, "-")
			lo, _ := strconv.ParseInt(part[:dash], 10, 64)
			hi, _ := strconv.ParseInt(part[dash+1:], 10, 64)
			ranges = append(ranges, rng{lo, hi})
		}
	}
	return ranges
}

func inRanges(n int64, ranges []rng) bool {
	for _, r := range ranges {
		if n >= r.lo && n <= r.hi {
			return true
		}
	}
	return false
}

func Part1(in Input) (res int) {
	ranges := parseRanges(in)

	for k := int64(1); k <= 5; k++ {
		pow := int64(1)
		for i := int64(0); i < k; i++ {
			pow *= 10
		}
		loS := pow / 10
		if k == 1 {
			loS = 1
		}
		for s := loS; s < pow; s++ {
			ss := s*pow + s
			if inRanges(ss, ranges) {
				res += int(ss)
			}
		}
	}

	return
}

func Part2(in Input) (res int) {
	ranges := parseRanges(in)

	// find max value across all ranges
	var maxVal int64
	for _, r := range ranges {
		if r.hi > maxVal {
			maxVal = r.hi
		}
	}

	// generate all invalid numbers: pattern s of length p repeated r>=2 times
	// deduplicate with a set (e.g. 111111 = "1"x6 = "11"x3 = "111"x2)
	seen := make(map[int64]bool)
	pow := [11]int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}

	for p := 1; p <= 5; p++ {
		loS := pow[p] / 10
		if p == 1 {
			loS = 1
		}
		for s := loS; s < pow[p]; s++ {
			for reps := 2; ; reps++ {
				// build s repeated reps times
				num := int64(0)
				for i := 0; i < reps; i++ {
					num = num*pow[p] + s
				}
				if num > maxVal {
					break
				}
				if !seen[num] && inRanges(num, ranges) {
					seen[num] = true
					res += int(num)
				}
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
