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

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	uf := NewUnionFind(n)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= maxDiff {
			uf.Union(i-1, i)
		}

	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]

		if uf.Find(u) == uf.Find(v) {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}