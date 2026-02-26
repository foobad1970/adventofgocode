# 2025 Day 6 - Trash Compactor

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.32     |
| Download input | 0.38     |
| LLM solve      | 110.22   |
| Compile        | 0.08     |
| Run            | 0.002    |

## Results

- Part 1: 5227286044585
- Part 2: 10227753257799

## Approach

**Part 1:** Parse the columnar worksheet by finding all-space separator columns, then grouping contiguous non-space columns into problems. For each group, extract the operator from the bottom row and numbers from each number row. Apply +/*, sum all results.

**Part 2:** Within each column group, each individual column holds one number — its digits are read top-to-bottom from the number rows (ignoring spaces). Process columns right-to-left within each group. Apply the operator to the resulting numbers and sum all grand totals.
