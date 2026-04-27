type Direction int

const (
	East Direction = iota
	South
	West
	North
)

var streets = map[int][2]Direction{
	1: {West, East},
	2: {North, South},
	3: {West, South},
	4: {East, South},
	5: {West, North},
	6: {East, North},
}

var moves = map[Direction][2]int{
	East:  {0, 1},
	South: {1, 0},
	West:  {0, -1},
	North: {-1, 0},
}

type Cell struct {
	r int
	c int
}

func opposite(dir Direction) Direction {
	switch dir {
	case East:
		return West
	case West:
		return East
	case North:
		return South
	case South:
		return North
	default:
		panic("invalid Direaction")
	}
}

func hasDirection(st [2]Direction, dir Direction) bool {
	return st[0] == dir || st[1] == dir
}

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	queue := []Cell{{0, 0}}
	visited[0][0] = true

	for head := 0; head < len(queue); head++ {
		cur := queue[head]

		if cur.r == m-1 && cur.c == n-1 {
			return true
		}

		st := streets[grid[cur.r][cur.c]]

		for _, dir := range st {
			move := moves[dir]

			nr := cur.r + move[0]
			nc := cur.c + move[1]

			if nr < 0 || nc < 0 || nr >= m || nc >= n {
				continue
			}

			if visited[nr][nc] {
				continue
			}

			nextStreet := streets[grid[nr][nc]]

			if !hasDirection(nextStreet, opposite(dir)) {
				continue
			}

			visited[nr][nc] = true
			queue = append(queue, Cell{nr, nc})

		}
	}

	return false
}