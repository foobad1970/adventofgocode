# Day 12 – Christmas Tree Farm

| Metric | Value |
|---|---|
| Read puzzle (s) | 0.33 |
| Download input (s) | 0.44 |
| LLM solve (s) | 702.38 |
| Compile (s) | 0.077 |
| Run (s) | 0.001 |

## Results

| Part | Answer |
|---|---|
| Part 1 | 476 |
| Part 2 | auto-star (no calculation) |

## Approach

### Part 1
2D bin-packing feasibility: given 6 polyomino shapes (each fitting in a 3×3
bounding box, with transparent `.` cells) and a W×H region, can all listed
presents be placed without `#`-cell overlap?

**Key insight**: all shapes fit in 3×3 bounding boxes. A region is feasible iff
`floor(W/3) × floor(H/3) ≥ sum(counts)` — i.e., there are enough non-overlapping
3×3 slots to assign one shape each. This is both necessary (each shape needs a
3×3 footprint) and sufficient (place each in its own slot).

For this input, the 3×3-slot condition exactly coincides with the simple area
condition `sum(count_i × area_i) ≤ W×H`, because the input was designed so that
the two conditions are equivalent for all 1000 regions.

Result: 524 infeasible (too many `#` cells for the grid), 476 feasible.

### Part 2
Auto-awarded star — story conclusion, no additional computation.
