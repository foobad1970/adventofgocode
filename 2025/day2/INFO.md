# 2025 Day 2 - Gift Shop

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.34     |
| Download input | 0.31     |
| LLM solve      | 124.54   |
| Compile        | 0.08     |
| Run            | 0.005    |

## Results

- Part 1: 19386344315
- Part 2: 34421651192

## Approach

**Part 1:** Generate all "invalid" IDs of the form `ss` (a digit string `s` concatenated with itself, no leading zeros), for half-lengths 1–5. For each candidate, check if it falls within any of the input ranges. Sum the hits.

**Part 2:** Extend to patterns repeated any number of times (≥2). Generate all candidates by trying every period length (1–5) and repetition count (2+) up to the maximum range value. Use a seen-set to deduplicate numbers that match multiple patterns (e.g. 111111 = "1"×6 = "11"×3 = "111"×2). Sum the unique hits.
