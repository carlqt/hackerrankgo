package hackerrankgo

import (
	"errors"
	"slices"
)

type Grid struct {
	Visited map[string]bool
	Matrix  [][]int32
}

func (g *Grid) get(x, y int) (int32, error) {
	if y < 0 || y >= len(g.Matrix) {
		return 0, errors.New("out of bounds")
	}
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

	// [1][2]
	return 0
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

		if len(queue) == 0 {
			break
		}
	}

	return int32(len(visited))
}

func pop(queue []point) (point, []point) {
	v := queue[0]
	result := slices.Delete(queue, 0, 1)

	return v, result
}

func adjacentFilledCell(start point, grid *Grid) []point {
	var adjacentFilled []point

	// check North - (-1, 0)
	// if grid.Matrix
}
