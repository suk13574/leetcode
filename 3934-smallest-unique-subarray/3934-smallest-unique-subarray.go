func smallestUniqueSubarray(nums []int) int {
	lo := 1
	hi := len(nums)

	for lo < hi {
		mid := lo + (hi-lo)/2

		if check(nums, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

type Hash struct {
	h1, h2 int
}

func check(nums []int, k int) bool {
	const (
		BASE = 100_001
		MOD1 = 1_000_000_007
		MOD2 = 1_000_000_009
	)

	sub := make(map[Hash]int)

	power1 := 1
	power2 := 1
	for i := 0; i < k-1; i++ {
		power1 = (power1 * BASE) % MOD1
		power2 = (power2 * BASE) % MOD2
	}

	h1, h2 := 0, 0
	for i, v := range nums {
		if i < k {
			h1 = (h1*BASE + v) % MOD1
			h2 = (h2*BASE + v) % MOD2

			if i == k-1 {
				sub[Hash{h1, h2}]++
			}
			continue
		}

		out := nums[i-k]
		h1 = (h1 - (out*power1)%MOD1 + MOD1) % MOD1
		h2 = (h2 - (out*power2)%MOD2 + MOD2) % MOD2

		h1 = (h1*BASE + v) % MOD1
		h2 = (h2*BASE + v) % MOD2

		sub[Hash{h1, h2}]++
	}

	for _, v := range sub {
		if v == 1 {
			return true
		}
	}

	return false
}
