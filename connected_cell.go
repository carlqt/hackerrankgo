package hackerrankgo

import (
	"errors"
	"slices"
)

type Grid struct {
	Visited []point
	Matrix  [][]int32
}

func (g *Grid) get(x, y int32) (int32, error) {
	if y < 0 || int(y) >= len(g.Matrix)-1 {
		return 0, errors.New("out of bounds")
	}

	if x < 0 || int(x) >= len(g.Matrix[0])-1 {
		return 0, errors.New("out of bounds")
	}

	return g.Matrix[y][x], nil
}

func (g *Grid) visited(x, y int32) bool {
	p := point{y, x}
	return slices.Contains(g.Visited, p)
}

type point [2]int32

func ConnectedCell(matrix [][]int32) int32 {
	var area int32

	grid := Grid{
		Matrix: matrix,
	}

	// Create the grid base on the matrix
	// 1,1,0,0
	// 1,0,1,0
	// 0,0,1,0
	// 1,0,0,0
	for i := range grid.Matrix {
		for i2, x := range grid.Matrix[i] {
			if x == 0 {
				continue
			}

			v, err := grid.get(int32(i2), int32(i))
			if err != nil || v == 0 {
				panic(err)
			}

			if grid.visited(int32(i2), int32(i)) {
				continue
			}

			start := point{int32(i2), int32(i)}
			result := bfs(start, &grid)
			if result > area {
				area = result
			}
		}
	}

	// [1][2]
	return area
}

func bfs(start point, grid *Grid) (area int32) {
	var visited []point
	var queue []point
	queue = append(queue, start)

	// Loop until the queue is empty
	// From the starting coordinate
	// get all adjacent coordinates and add them to the queue
	// Add the current coordinate to visited
	for {
		// Remove the first element from the queue
		current, queue := pop(queue)
		visited = append(visited, current)

		// Get the adjacent coordinates
		// avoid panicing for index out of bounds
		// add the adjacent coordinates to the queue
		adjacentCells := adjacentFilledCell(current, grid)
		for _, a := range adjacentCells {
			if !slices.Contains(visited, a) || !slices.Contains(queue, a) {
				queue = append(queue, a)
			}
		}

		if len(queue) == 0 {
			break
		}
	}

	// Add visited to grid.visited
	for _, v := range visited {
		if !slices.Contains(grid.Visited, v) {
			grid.Visited = append(grid.Visited, v)
		}
	}

	return int32(len(visited))
}

func pop(queue []point) (point, []point) {
	v := queue[0]
	result := slices.Delete(queue, 0, 1)

	return v, result
}

// This should not return the visited nodes
func adjacentFilledCell(start point, grid *Grid) []point {
	var adjacentFilled []point

	// Check coordinate north of starting point (-1, 0)
	p := point{start[1] - 1, start[0]}
	r, err := grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate north east of starting point (-1, +1)
	p = point{start[1] - 1, start[0] + 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate east of starting point (0, +1)
	p = point{start[1], start[0] + 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate south-east of starting point (+1, +1)
	p = point{start[1] + 1, start[0] + 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate south of starting point (+1, 0)
	p = point{start[1] + 1, start[0]}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate south-west of starting point (+1, -1)
	p = point{start[1] + 1, start[0] - 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate west of starting point (0, -1)
	p = point{start[1], start[0] - 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	// Check coordinate north-west of starting point (-1, -1)
	p = point{start[1] - 1, start[0] - 1}
	r, err = grid.get(p[1], p[0])
	if err == nil && r == 1 {
		adjacentFilled = append(adjacentFilled, p)
	}

	return adjacentFilled
}
