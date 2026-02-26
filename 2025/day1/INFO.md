# 2025 Day 1 - Secret Entrance

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.37     |
| Download input | 0.34     |
| LLM solve      | 75.29    |
| Compile        | 0.09     |
| Run            | 0.003    |

## Results

- Part 1: 999
- Part 2: 6099

## Approach

**Part 1:** Simulate the dial (0-99, wrapping) starting at 50. For each rotation (L/R + amount), compute the new position using modular arithmetic. Count how many times the final position is exactly 0.

**Part 2:** Count every individual click that lands on 0 — not just the final position. Solved mathematically: given a rotation of `a` clicks from position `pos`, compute the first click that hits 0 (the "offset"), then count how many multiples of 100 fit in the remaining clicks.
