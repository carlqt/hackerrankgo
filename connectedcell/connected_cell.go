package connectedcell

type Cell struct {
	y, x int32
}

type Grid [][]int32

func (g Grid) get(x, y int32) int32 {
	if y < 0 || int(y) >= len(g) {
		return -1
	}

	if x < 0 || int(x) >= len(g[0]) {
		return -1
	}

	return g[y][x]
}

func connectedCell(matrix [][]int32) int32 {
	var visitedCells []Cell
	var area int

	grid := Grid(matrix)

	for i, y := range grid {
		for i2, x := range y {
			cell := Cell{int32(i), int32(i2)}

			if x <= 0 {
				continue
			}

			if Contains(visitedCells, cell) {
				continue
			}

			filledCells := bfs(cell, grid)
			if len(filledCells) > area {
				area = len(filledCells)
			}

			visitedCells = append(visitedCells, filledCells...)
		}
	}

	return int32(area)
}

func bfs(start Cell, grid Grid) []Cell {
	var visited []Cell
	var queue []Cell

	queue = append(queue, start)

	for len(queue) > 0 {
		current := pop(&queue)
		visited = append(visited, current)

		for _, a := range adjacentFilledCell(current, grid) {
			if !Contains(visited, a) && !Contains(queue, a) {
				queue = append(queue, a)
			}
		}
	}

	return visited
}

func pop(queue *[]Cell) Cell {
	v := (*queue)[0]
	*queue = (*queue)[1:]

	return v
}

func Contains(Cells []Cell, p Cell) bool {
	for _, pp := range Cells {
		if pp.x == p.x && pp.y == p.y {
			return true
		}
	}

	return false
}

func adjacentFilledCell(start Cell, grid Grid) []Cell {
	var adjacentFilled []Cell

	north := Cell{-1, 0}
	northEast := Cell{-1, 1}
	east := Cell{0, 1}
	southEast := Cell{1, 1}
	south := Cell{1, 0}
	southWest := Cell{1, -1}
	west := Cell{0, -1}
	northWest := Cell{-1, -1}

	orientations := []Cell{
		north,
		northEast,
		east,
		southEast,
		south,
		southWest,
		west,
		northWest,
	}

	for _, p := range orientations {
		adjacentCell := Cell{
			y: start.y + p.y,
			x: start.x + p.x,
		}

		value := grid.get(adjacentCell.x, adjacentCell.y)
		if value == 1 {
			adjacentFilled = append(adjacentFilled, adjacentCell)
		}
	}

	return adjacentFilled
}
