# Day 10 – Factory

| Metric | Value |
|---|---|
| Read puzzle (s) | 0.32 |
| Download input (s) | 0.31 |
| LLM solve (s) | 1646.85 |
| Compile (s) | 0.09 |
| Run (s) | 0.177 |

## Results

| Part | Answer |
|---|---|
| Part 1 | 502 |
| Part 2 | 21467 |

## Approach

### Part 1
GF(2) (binary XOR) linear algebra. Each button toggles a bitmask of lights.
Brute-force all 2^m subsets (m ≤ 13, so at most 8192 per machine).
Find the minimum-weight subset whose XOR equals the target pattern.
Sum across 198 machines.

### Part 2
Buttons now *increment* integer counters. Minimize total button presses.
This is a non-negative Integer Linear Program (ILP): minimize sum(x) subject to Ax = b, x ≥ 0, x integer.

**Algorithm: Null space parametrization + vertex enumeration**
1. Gaussian elimination (exact rational arithmetic via big.Rat) → particular
   solution x_p and null space basis N (dimension d = 0..3).
2. General solution: x = x_p + N·z, z ∈ Z^d, z ≥ 0.
3. Enumerate all C(m+d, d) vertex candidates of the LP relaxation: for each
   d-subset of constraints (type A: x_i=0, type B: z_k=0), solve the d×d
   system and check feasibility. Collect tight z-bounds from feasible vertices.
4. Enumerate integer z in computed bounds; for each, verify x ≥ 0 (integer)
   and track minimum sum.
