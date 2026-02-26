# 2025 Day 3 - Lobby

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.31     |
| Download input | 0.36     |
| LLM solve      | 110.51   |
| Compile        | 0.12     |
| Run            | 0.002    |

## Results

- Part 1: 17031
- Part 2: 168575096286051

## Approach

**Part 1:** For each bank (line of digits), find the maximum 2-digit number formable by picking two digits in order. Precompute the max digit to the right of each position, then scan left-to-right choosing the best pair in O(n).

**Part 2:** Generalized to picking exactly 12 digits. Classic greedy "max k digits from string" algorithm: at each selection step, pick the largest digit in the allowable window (adjusted so enough digits remain), advancing the start pointer past the chosen position.
