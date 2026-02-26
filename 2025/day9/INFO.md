# 2025 Day 9 - Movie Theater

## Timings

| Step           | Time (s) |
|----------------|----------|
| Read puzzle    | 0.33     |
| Download input | 0.29     |
| LLM solve      | 230.70   |
| Compile        | 0.09     |
| Run            | 0.014    |

## Results

- Part 1: 4763932976
- Part 2: 1501292304

## Approach

**Part 1:** Brute force O(n²) over all pairs of red tiles. Area of rectangle = (|dx|+1) × (|dy|+1) since corners are inclusive. 496 tiles → ~123k pairs, trivially fast.

**Part 2:** The red tiles form a rectilinear polygon; green tiles fill the edges and interior. A valid rectangle must lie entirely within this region. Used coordinate compression (248 unique x, 248 unique y values) to create a ~61k-cell compressed grid. For each interior gap cell (between consecutive unique x and y values), determined inside/outside using ray casting with doubled coordinates to avoid fractions. Built a 2D prefix sum over the compressed grid. For each pair of red tiles as corners, verified all gap cells within are inside (O(1) per pair via prefix sum), then maximized area.
