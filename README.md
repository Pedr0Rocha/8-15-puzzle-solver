# 8-Puzzle & 15-Puzzle Solver

## 8-Puzzle solvers

- DFS (brute force)
- A Star [Manhattan distance]
- A Star using Priority Queue [Manhattan distance]

### DFS - Brute Force

DFS will expand nodes until it finds the board solution. Since the max
number of combinations it not too high for 8-Puzzle 3x3 (around 180k)
it runs fine.

### A Star [Manhattan distance]

A\* will search for the minimum amount of steps to reach the solution
board. The heuristic used here is Manhattan distance. This process can
take quite a while for a 31 steps board, which is the maximum amount
of steps to get to a solution.

31 steps board: ~6s

### A Star using Priority Queue [Manhattan distance]

This version is the same as the previous one, but instead of using a
hashmap to manage the fScores, it uses a PriorityQueue. This is a huge
improvement to the algorithm since we need to calculate the lowest
fScore every time we process a node.

31 steps board: ~0.15s

## 15-Puzzle solvers
