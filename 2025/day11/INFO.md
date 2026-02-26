# Day 11 – Reactor

| Metric | Value |
|---|---|
| Read puzzle (s) | 8.39 |
| Download input (s) | 0.31 |
| LLM solve (s) | 128.94 |
| Compile (s) | 0.073 |
| Run (s) | 0.001 |

## Results

| Part | Answer |
|---|---|
| Part 1 | 696 |
| Part 2 | 473741288064360 |

## Approach

### Part 1
Directed acyclic graph (DAG) path count from "you" to "out".
Memoized DFS: `count(node)` = sum of `count(child)` for each child.
Base case: `count("out") = 1`. 596 nodes, answer in <1ms.

### Part 2
Count paths from "svr" to "out" that visit both "dac" and "fft" (any order).
Extended DP with a 2-bit visited mask (bit 0 = dac visited, bit 1 = fft visited).
State: (current_node, mask). Transition: entering child c updates mask |= bit(c).
At "out": count 1 only if mask == 3 (both visited). O(nodes × 4) states.
