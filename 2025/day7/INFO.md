# 2025 Day 7 - Laboratories

## Timings

| Step           | Time (s)          |
|----------------|-------------------|
| Read puzzle    | 0.31              |
| Download input | 0.39              |
| LLM solve      | 169.30 (includes user-requested pause) |
| Compile        | 0.08              |
| Run            | 0.003             |

## Results

- Part 1: 1681
- Part 2: 422102272495018

## Approach

**Part 1:** BFS simulation of downward-moving beams. Start at `S`, move down each step. On hitting `^`, increment split count and enqueue beams at the left and right neighbors (same row). Track visited positions to avoid reprocessing. Total unique `^` cells reached = total splits.

**Part 2:** Memoized recursion — `timelines(c, r)` returns the number of exit paths from that position. At `.`/`S`: recurse one row down. At `^`: sum `timelines(c-1, r) + timelines(c+1, r)`. Out-of-bounds laterally = 0 timelines; off the bottom = 1 (particle exits). Confirmed no adjacent horizontal splitters, so no cycles possible.
