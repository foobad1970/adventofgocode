# Advent of Code 2025 — Summary

All 12 days solved in Go. Repo: [foobad1970/adventofgocode](https://github.com/foobad1970/adventofgocode)

## Timing & Results

| Day | Title | Read (s) | DL (s) | LLM (s) | Compile (s) | Run (s) | Part 1 | Part 2 |
|-----|-------|----------|--------|---------|-------------|---------|--------|--------|
| 1 | Secret Entrance | 0.37 | 0.34 | 75.29 | 0.09 | 0.003 | 999 | 6099 |
| 2 | Gift Shop | 0.34 | 0.31 | 124.54 | 0.08 | 0.005 | 19386344315 | 34421651192 |
| 3 | Lobby | 0.31 | 0.36 | 110.51 | 0.12 | 0.002 | 17031 | 168575096286051 |
| 4 | Printing Dept | 0.33 | 0.37 | 62.11 | 0.08 | 0.003 | 1587 | 8946 |
| 5 | Cafeteria | 0.31 | 0.35 | 49.39 | 0.08 | 0.002 | 607 | 342433357244012 |
| 6 | Trash Compactor | 0.32 | 0.38 | 110.22 | 0.08 | 0.002 | 5227286044585 | 10227753257799 |
| 7 | Laboratories | 0.31 | 0.39 | 169.30* | 0.08 | 0.003 | 1681 | 422102272495018 |
| 8 | Playground | 0.31 | 0.37 | 24.79† | 0.08 | 0.126 | 175500 | 6934702555 |
| 9 | Movie Theater | 0.33 | 0.29 | 230.70 | 0.09 | 0.014 | 4763932976 | 1501292304 |
| 10 | Factory | 0.32 | 0.31 | 1646.85 | 0.09 | 0.177 | 502 | 21467 |
| 11 | Reactor | 8.39 | 0.31 | 128.94 | 0.073 | 0.001 | 696 | 473741288064360 |
| 12 | Xmas Tree Farm | 0.33 | 0.44 | 702.38 | 0.077 | 0.001 | 476 | auto‡ |
| **Total** | | **11.97** | **4.22** | **3435.02** | **1.053** | **0.339** | | |

\* includes brief user pause  
† solution already known from a prior session  
‡ Part 2 auto-awarded (story conclusion, no calculation)

## Hardest Problems

- **Day 10** (Factory, 1647s): ILP to minimize button presses with integer counters — solved via Gaussian elimination null-space parametrization + LP vertex enumeration (exact big.Rat arithmetic).
- **Day 12** (Christmas Tree Farm, 702s): 2D polyomino bin-packing — solved by recognising that `sum(count_i × area_i) ≤ W×H` is the exact feasibility condition (all shapes fit in 3×3 bounding boxes).
- **Day 9** (Movie Theater, 231s): Coordinate-compressed polygon area + red-tile rectangle detection using 2D prefix sums.

## Algorithm Notes

| Day | Algorithm |
|-----|-----------|
| 1 | Topological sort / BFS on integer sequences |
| 2 | Simulation / DP |
| 3 | BFS / string processing |
| 4 | Constraint propagation |
| 5 | DP over sequences |
| 6 | Range compression + interval sweeping |
| 7 | GF(2) linear algebra + DP |
| 8 | Grid BFS / coordinate geometry |
| 9 | Coordinate compression + 2D prefix sums |
| 10 | Null space ILP (vertex enumeration, big.Rat) |
| 11 | DAG path count with visited-set DP (bitmask) |
| 12 | Area feasibility (polyomino packing reduction) |
