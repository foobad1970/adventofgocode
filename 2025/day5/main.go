package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rng struct{ lo, hi int64 }

type Input struct {
	ranges []rng
	ids    []int64
}

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			dash := strings.Index(line, "-")
			lo, _ := strconv.ParseInt(line[:dash], 10, 64)
			hi, _ := strconv.ParseInt(line[dash+1:], 10, 64)
			res.ranges = append(res.ranges, rng{lo, hi})
		} else {
			id, _ := strconv.ParseInt(line, 10, 64)
			res.ids = append(res.ids, id)
		}
	}
	// sort and merge ranges
	sort.Slice(res.ranges, func(i, j int) bool {
		return res.ranges[i].lo < res.ranges[j].lo
	})
	merged := []rng{res.ranges[0]}
	for _, r := range res.ranges[1:] {
		last := &merged[len(merged)-1]
		if r.lo <= last.hi+1 {
			if r.hi > last.hi {
				last.hi = r.hi
			}
		} else {
			merged = append(merged, r)
		}
	}
	res.ranges = merged
	return
}

func isFresh(id int64, ranges []rng) bool {
	lo, hi := 0, len(ranges)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if id < ranges[mid].lo {
			hi = mid - 1
		} else if id > ranges[mid].hi {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func Part1(in Input) (res int) {
	for _, id := range in.ids {
		if isFresh(id, in.ranges) {
			res++
		}
	}
	return
}

func Part2(in Input) (res int) {
	for _, r := range in.ranges {
		res += int(r.hi - r.lo + 1)
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
