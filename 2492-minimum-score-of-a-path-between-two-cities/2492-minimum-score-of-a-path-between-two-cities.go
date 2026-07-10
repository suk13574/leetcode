type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	if uf.size[rootA] < uf.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	uf.parent[rootB] = rootA
	uf.size[rootA] += uf.size[rootB]

	return true
}

func minScore(n int, roads [][]int) int {
	uf := NewUnionFind(n + 1)

	for _, road := range roads {
		u, v := road[0], road[1]

		uf.Union(u, v)
	}

	root := uf.Find(1)
	res := 1 << 30

	for _, road := range roads {
		u, d := road[0], road[2]

		if uf.Find(u) == root {
			res = min(res, d)
		}
	}

	return res
}