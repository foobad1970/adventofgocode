package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"math/big"
	"math/bits"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type machine struct {
	// Part 1
	p1target  uint32
	p1buttons []uint32
	// Part 2
	p2targets []int
	p2buttons [][]int
}

type Input []machine

var reLights = regexp.MustCompile(`\[([.#]+)\]`)
var reButton = regexp.MustCompile(`\(([^)]+)\)`)
var reTargets = regexp.MustCompile(`\{([^}]+)\}`)

func parseButton1(s string) uint32 {
	var mask uint32
	for _, part := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(part))
		mask |= 1 << uint(n)
	}
	return mask
}

func parseButton2(s string) []int {
	var indices []int
	for _, part := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(part))
		indices = append(indices, n)
	}
	return indices
}

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		m := machine{}
		lm := reLights.FindStringSubmatch(line)
		if lm != nil {
			for i, ch := range lm[1] {
				if ch == '#' {
					m.p1target |= 1 << uint(i)
				}
			}
		}
		for _, bm := range reButton.FindAllStringSubmatch(line, -1) {
			m.p1buttons = append(m.p1buttons, parseButton1(bm[1]))
			m.p2buttons = append(m.p2buttons, parseButton2(bm[1]))
		}
		tm := reTargets.FindStringSubmatch(line)
		if tm != nil {
			for _, part := range strings.Split(tm[1], ",") {
				n, _ := strconv.Atoi(strings.TrimSpace(part))
				m.p2targets = append(m.p2targets, n)
			}
		}
		res = append(res, m)
	}
	return
}

// Part 1: GF(2) brute force
func minPresses1(m machine) int {
	nb := len(m.p1buttons)
	best := nb + 1
	for mask := 0; mask < (1 << nb); mask++ {
		var xorVal uint32
		for i := 0; i < nb; i++ {
			if mask&(1<<i) != 0 {
				xorVal ^= m.p1buttons[i]
			}
		}
		if xorVal == m.p1target {
			w := bits.OnesCount(uint(mask))
			if w < best {
				best = w
			}
		}
	}
	if best == nb+1 {
		return 0
	}
	return best
}

func Part1(in Input) (res int) {
	for _, m := range in {
		res += minPresses1(m)
	}
	return
}

// ---------- Part 2: ILP via null space + vertex enumeration ----------

// gaussJordan returns particular solution xp and null space basis vectors.
// Returns nil, nil if system is inconsistent.
func gaussJordan(A [][]int, b []int) (xp []*big.Rat, nullBasis [][]*big.Rat) {
	rows := len(A)
	if rows == 0 {
		return nil, nil
	}
	cols := len(A[0])

	mat := make([][]*big.Rat, rows)
	for i := range mat {
		mat[i] = make([]*big.Rat, cols+1)
		for j := 0; j < cols; j++ {
			mat[i][j] = big.NewRat(int64(A[i][j]), 1)
		}
		mat[i][cols] = big.NewRat(int64(b[i]), 1)
	}

	pivotCols := []int{}
	row := 0
	for col := 0; col < cols && row < rows; col++ {
		pivotRow := -1
		for r := row; r < rows; r++ {
			if mat[r][col].Sign() != 0 {
				pivotRow = r
				break
			}
		}
		if pivotRow == -1 {
			continue
		}
		mat[row], mat[pivotRow] = mat[pivotRow], mat[row]
		f := new(big.Rat).Set(mat[row][col])
		for j := 0; j <= cols; j++ {
			mat[row][j].Quo(mat[row][j], f)
		}
		for r := 0; r < rows; r++ {
			if r != row && mat[r][col].Sign() != 0 {
				g := new(big.Rat).Set(mat[r][col])
				for j := 0; j <= cols; j++ {
					tmp := new(big.Rat).Mul(g, mat[row][j])
					mat[r][j].Sub(mat[r][j], tmp)
				}
			}
		}
		pivotCols = append(pivotCols, col)
		row++
	}

	for r := row; r < rows; r++ {
		if mat[r][cols].Sign() != 0 {
			return nil, nil
		}
	}

	pivotSet := make(map[int]bool)
	for _, c := range pivotCols {
		pivotSet[c] = true
	}
	freeCols := []int{}
	for c := 0; c < cols; c++ {
		if !pivotSet[c] {
			freeCols = append(freeCols, c)
		}
	}

	xp = make([]*big.Rat, cols)
	for j := range xp {
		xp[j] = new(big.Rat)
	}
	for i, c := range pivotCols {
		xp[c] = new(big.Rat).Set(mat[i][cols])
	}

	nullBasis = make([][]*big.Rat, len(freeCols))
	for k, fc := range freeCols {
		vec := make([]*big.Rat, cols)
		for j := range vec {
			vec[j] = new(big.Rat)
		}
		vec[fc] = big.NewRat(1, 1)
		for i, c := range pivotCols {
			vec[c] = new(big.Rat).Neg(mat[i][fc])
		}
		nullBasis[k] = vec
	}
	return
}

// solveSquare solves the d×d system M*z = rhs using Gaussian elimination.
// Returns nil if singular.
func solveSquare(M [][]*big.Rat, rhs []*big.Rat) []*big.Rat {
	n := len(M)
	if n == 0 {
		return []*big.Rat{}
	}
	mat := make([][]*big.Rat, n)
	for i := range mat {
		mat[i] = make([]*big.Rat, n+1)
		for j := 0; j < n; j++ {
			mat[i][j] = new(big.Rat).Set(M[i][j])
		}
		mat[i][n] = new(big.Rat).Set(rhs[i])
	}
	for col := 0; col < n; col++ {
		p := -1
		for r := col; r < n; r++ {
			if mat[r][col].Sign() != 0 {
				p = r
				break
			}
		}
		if p == -1 {
			return nil
		}
		mat[col], mat[p] = mat[p], mat[col]
		f := new(big.Rat).Set(mat[col][col])
		for j := 0; j <= n; j++ {
			mat[col][j].Quo(mat[col][j], f)
		}
		for r := 0; r < n; r++ {
			if r != col && mat[r][col].Sign() != 0 {
				g := new(big.Rat).Set(mat[r][col])
				for j := 0; j <= n; j++ {
					tmp := new(big.Rat).Mul(g, mat[col][j])
					mat[r][j].Sub(mat[r][j], tmp)
				}
			}
		}
	}
	z := make([]*big.Rat, n)
	for i := range z {
		z[i] = new(big.Rat).Set(mat[i][n])
	}
	return z
}

func ratFloor(r *big.Rat) int64 {
	f, _ := r.Float64()
	return int64(math.Floor(f))
}

func ratCeil(r *big.Rat) int64 {
	f, _ := r.Float64()
	return int64(math.Ceil(f))
}

func minPressesPart2(m machine) int {
	nc := len(m.p2targets)
	nb := len(m.p2buttons)
	if nb == 0 || nc == 0 {
		return 0
	}

	A := make([][]int, nc)
	for i := range A {
		A[i] = make([]int, nb)
	}
	for j, btn := range m.p2buttons {
		for _, idx := range btn {
			A[idx][j] = 1
		}
	}

	xp, null := gaussJordan(A, m.p2targets)
	if xp == nil {
		return -1
	}
	d := len(null)

	sumRat := func(zVals []*big.Rat) (int, bool) {
		total := 0
		for i := 0; i < nb; i++ {
			xi := new(big.Rat).Set(xp[i])
			for k := 0; k < d; k++ {
				tmp := new(big.Rat).Mul(zVals[k], null[k][i])
				xi.Add(xi, tmp)
			}
			if !xi.IsInt() || xi.Sign() < 0 {
				return 0, false
			}
			total += int(xi.Num().Int64())
		}
		return total, true
	}

	if d == 0 {
		zz := []*big.Rat{}
		if t, ok := sumRat(zz); ok {
			return t
		}
		return -1
	}

	// Constraints: nb type-A (x_p[i]+N*z >= 0), d type-B (z_k >= 0)
	totalC := nb + d
	zLo := make([]*big.Rat, d)
	zHi := make([]*big.Rat, d)
	for k := 0; k < d; k++ {
		zLo[k] = new(big.Rat)
		zHi[k] = new(big.Rat)
	}
	eps := big.NewRat(-1, 1000000)

	chosen := make([]int, d)
	var vertexEnum func(start, k int)
	vertexEnum = func(start, k int) {
		if k == d {
			M := make([][]*big.Rat, d)
			rhs := make([]*big.Rat, d)
			for i := 0; i < d; i++ {
				c := chosen[i]
				M[i] = make([]*big.Rat, d)
				if c < nb {
					for k2 := 0; k2 < d; k2++ {
						M[i][k2] = new(big.Rat).Set(null[k2][c])
					}
					rhs[i] = new(big.Rat).Neg(xp[c])
				} else {
					for k2 := 0; k2 < d; k2++ {
						M[i][k2] = new(big.Rat)
					}
					M[i][c-nb] = big.NewRat(1, 1)
					rhs[i] = new(big.Rat)
				}
			}
			z := solveSquare(M, rhs)
			if z == nil {
				return
			}
			for k2 := 0; k2 < d; k2++ {
				if z[k2].Cmp(eps) < 0 {
					return
				}
			}
			for i := 0; i < nb; i++ {
				xi := new(big.Rat).Set(xp[i])
				for k2 := 0; k2 < d; k2++ {
					tmp := new(big.Rat).Mul(z[k2], null[k2][i])
					xi.Add(xi, tmp)
				}
				if xi.Cmp(eps) < 0 {
					return
				}
			}
			for k2 := 0; k2 < d; k2++ {
				if z[k2].Cmp(zLo[k2]) < 0 {
					zLo[k2] = new(big.Rat).Set(z[k2])
				}
				if z[k2].Cmp(zHi[k2]) > 0 {
					zHi[k2] = new(big.Rat).Set(z[k2])
				}
			}
			return
		}
		for i := start; i <= totalC-(d-k); i++ {
			chosen[k] = i
			vertexEnum(i+1, k+1)
		}
	}
	vertexEnum(0, 0)

	loInts := make([]int64, d)
	hiInts := make([]int64, d)
	for k := 0; k < d; k++ {
		lo := ratFloor(zLo[k])
		if lo < 0 {
			lo = 0
		}
		loInts[k] = lo
		hiInts[k] = ratCeil(zHi[k]) + 1
	}

	best := -1
	zCur := make([]*big.Rat, d)
	for k := range zCur {
		zCur[k] = new(big.Rat)
	}

	var search func(k int)
	search = func(k int) {
		if k == d {
			if t, ok := sumRat(zCur); ok {
				if best < 0 || t < best {
					best = t
				}
			}
			return
		}
		for z := loInts[k]; z <= hiInts[k]; z++ {
			zCur[k] = big.NewRat(z, 1)
			search(k + 1)
		}
	}
	search(0)
	return best
}

func Part2(in Input) (res int) {
	for _, m := range in {
		v := minPressesPart2(m)
		if v < 0 {
			log.Printf("warning: no solution for machine")
			continue
		}
		res += v
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
