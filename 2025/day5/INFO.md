# 2025 Day 5 - Cafeteria

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.31     |
| Download input | 0.35     |
| LLM solve      | 49.39    |
| Compile        | 0.08     |
| Run            | 0.002    |

## Results

- Part 1: 607
- Part 2: 342433357244012

## Approach

**Part 1:** Parse fresh ID ranges and ingredient IDs (int64 — values up to ~538 trillion). Sort and merge overlapping ranges, then binary-search each ingredient ID against the merged set. Count fresh ones.

**Part 2:** The answer is simply the total number of integers covered by the merged ranges — sum `(hi - lo + 1)` for each merged range.
