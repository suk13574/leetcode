func xorAfterQueries(nums []int, queries [][]int) int {
	const MOD int64 = 1_000_000_007
	n := len(nums)
	B := int(math.Sqrt(float64(n))) + 1 // square root

	mul := make([]int64, n)
	for i := 0; i < n; i++ {
		mul[i] = 1
	}

	small := make([][][]int64, B+1)
	for k := 1; k <= B; k++ {
		small[k] = make([][]int64, k)
		for r := 0; r < k; r++ {
			size := 0
			if r < n {
				size = (n-1-r)/k + 1
			}
			small[k][r] = make([]int64, size+1)
			for i := range small[k][r] {
				small[k][r][i] = 1
			}
		}
	}

	// a^e % MOD
	powMod := func(a, e int64) int64 {
		res := int64(1)
		for e > 0 {
			if e&1 == 1 {
				res = res * a % MOD
			}
			a = a * a % MOD
			e >>= 1
		}
		return res
	}

	// modular inverse of x
	invMod := func(x int64) int64 {
		return powMod(x, MOD-2)
	}

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], int64(q[3])

		if k > B {
			for i := l; i <= r; i += k {
				mul[i] = mul[i] * v % MOD
			}
		} else {
			rem := l % k
			L := l / k
			R := (r - rem) / k

			diff := small[k][rem]
			diff[L] = diff[L] * v % MOD
			diff[R+1] = diff[R+1] * invMod(v) % MOD
		}
	}

	for k := 1; k <= B; k++ {
		for r := 0; r < k; r++ {
			diff := small[k][r]
			cur := int64(1)
			idx, pos := r, 0

			for idx < n {
				cur = cur * diff[pos] % MOD
				mul[idx] = mul[idx] * cur % MOD
				idx += k
				pos++
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		val := int64(nums[i]) * mul[i] % MOD
		res ^= int(val)
	}

	return res
}