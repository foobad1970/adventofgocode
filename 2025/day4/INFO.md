# 2025 Day 4 - Printing Department

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.33     |
| Download input | 0.37     |
| LLM solve      | 62.11    |
| Compile        | 0.08     |
| Run            | 0.003    |

## Results

- Part 1: 1587
- Part 2: 8946

## Approach

**Part 1:** For each `@` cell, count its 8-directional neighbors. If fewer than 4 are also `@`, the roll is accessible. Sum the count.

**Part 2:** BFS peel: seed the queue with all initially accessible rolls. When a roll is removed, decrement its neighbors' effective counts — any neighbor now with < 4 `@` neighbors gets enqueued. Repeat until the queue is empty. Total removed is the answer.
