# 2025 Day 8 - Playground

## Timings

| Step           | Time (s)                               |
|----------------|----------------------------------------|
| Read puzzle    | 0.31                                   |
| Download input | 0.37                                   |
| LLM solve      | 24.79 (solution already known from prior run) |
| Compile        | 0.08                                   |
| Run            | 0.126                                  |

## Results

- Part 1: 175500
- Part 2: 6934702555

## Approach

**Part 1:** Compute all pairwise squared Euclidean distances between 1000 3D junction boxes (499,500 pairs). Sort by distance, connect the 1000 closest pairs using Union-Find. Tally circuit sizes by root, sort descending, multiply the top 3.

**Part 2:** Continue Kruskal's MST beyond 1000 pairs — keep unioning until all nodes are in one component. The last union that completes the spanning tree gives the final two junction boxes; multiply their X coordinates.

*Note: this puzzle was solved in a prior session; the LLM solve time reflects transcription speed, not original reasoning.*
