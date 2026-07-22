import "sort"

type ZeroSection struct {
	left   int
	right  int
	length int
}

type MaxSegmentTree struct {
	n    int
	tree []int
}

func NewMaxSegmentTree(values []int) *MaxSegmentTree {
	n := len(values)
	tree := make([]int, 2*n)

	// leaf node
	copy(tree[n:], values)

	// parent node
	for i := n - 1; i > 0; i-- {
		tree[i] = max(tree[i*2], tree[i*2+1])
	}

	return &MaxSegmentTree{
		n:    n,
		tree: tree,
	}
}

func (st *MaxSegmentTree) Query(left, right int) int {
	if st.n == 0 || left > right {
		return 0
	}

	left += st.n
	right += st.n

	res := 0

	for left <= right {
		if left%2 == 1 {
			res = max(res, st.tree[left])
			left++
		}

		if right%2 == 0 {
			res = max(res, st.tree[right])
			right--
		}

		left /= 2
		right /= 2
	}

	return res
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)

	zss := make([]ZeroSection, 0)
	totalOne := 0
	left := -1
	for i := 0; i <= n; i++ {
		ch := byte('1')
		if i < n {
			ch = s[i]
		}

		if ch == '0' {
			if left == -1 {
				left = i
			}
			continue
		}

		if i < n {
			totalOne++
		}

		if left != -1 {
			right := i - 1
			zss = append(zss, ZeroSection{
				left:   left,
				right:  right,
				length: right - left + 1,
			})

			left = -1
		}
	}

	gains := make([]int, max(0, len(zss)-1))
	for i := 0; i+1 < len(zss); i++ {
		gains[i] = zss[i].length + zss[i+1].length
	}

	st := NewMaxSegmentTree(gains)

	res := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]

		res[i] = totalOne

		first := sort.Search(len(zss), func(i int) bool {
			return zss[i].right >= l
		})

		end := sort.Search(len(zss), func(i int) bool {
			return zss[i].left > r
		})

		if end-first < 2 {
			continue
		}

		last := end - 1

		overlapLen := func(zs ZeroSection) int {
			return min(zs.right, r) - max(zs.left, l) + 1
		}

		leftGain := overlapLen(zss[first]) + overlapLen(zss[first+1])
		rightGain := overlapLen(zss[last-1]) + overlapLen(zss[last])

		bestGain := max(leftGain, rightGain)

		if last-first >= 3 {
			internalGain := st.Query(first+1, last-2)
			bestGain = max(bestGain, internalGain)
		}

		res[i] += bestGain
	}

	return res
}