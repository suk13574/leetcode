type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{
		parent: p,
	}
}

func (d *DSU) Find(a int) int {
	if d.parent[a] != a {
		d.parent[a] = d.Find(d.parent[a])
	}
	return d.parent[a]
}

func (d *DSU) Union(a, b int) bool {
	ap, bp := d.Find(a), d.Find(b)

	if ap != bp {
		d.parent[bp] = ap
		return true
	}

	return false
}

func maxStability(n int, edges [][]int, k int) int {
	can := func(x int) bool {
		dsu := NewDSU(n)
		need := n

		freeEdges := make([][]int, 0)
		upEdges := make([][]int, 0)

		for _, e := range edges {
			u, v, s, must := e[0], e[1], e[2], e[3]

			if must == 1 {
				if s < x {
					return false
				}
				if !dsu.Union(u, v) {
					return false
				}
				need--
			} else {
				if s >= x {
					freeEdges = append(freeEdges, []int{u, v})
				} else if s*2 >= x {
					upEdges = append(upEdges, []int{u, v})
				}
			}
		}

		for _, e := range freeEdges {
			if dsu.Union(e[0], e[1]) {
				need--
			}
		}

		usedUpgrade := 0
		for _, e := range upEdges {
			if dsu.Union(e[0], e[1]) {
				usedUpgrade++
				if usedUpgrade > k {
					return false
				}
				need--
			}
		}

		return need == 1
	}

	hi := 0
	for _, e := range edges {
		s, must := e[2], e[3]
		if must == 0 {
			s *= 2
		}
		hi = max(hi, s)
	}

	lo := 1
	res := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if can(mid) {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return res
}