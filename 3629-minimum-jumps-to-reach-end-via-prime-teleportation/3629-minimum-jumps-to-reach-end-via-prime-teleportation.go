type State struct {
	idx  int
	step int
}

func minJumps(nums []int) int {
	n := len(nums)

	maxVal := 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}

	spf := buildSPF(maxVal)

	isPrime := func(x int) bool {
		return x >= 2 && spf[x] == x
	}

	primeToIndices := make(map[int][]int)

	for i, num := range nums {
		factors := getUniquePrimeFactors(num, spf)

		for _, f := range factors {
			primeToIndices[f] = append(primeToIndices[f], i)
		}
	}

	visited := make([]bool, n)
	visited[0] = true

	queue := []State{{idx: 0, step: 0}}

	for head := 0; head < len(queue); head++ {
		cur := queue[head]

		if cur.idx == n-1 {
			return cur.step
		}

		for _, move := range []int{-1, 1} {
			next := cur.idx + move
			if next >= 0 && next < n && !visited[next] {
				visited[next] = true
				queue = append(queue, State{next, cur.step + 1})

			}
		}

		num := nums[cur.idx]
		if isPrime(num) {
			for _, next := range primeToIndices[num] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, State{next, cur.step + 1})
				}
			}
		}

		primeToIndices[num] = nil
	}

	return -1
}

func buildSPF(maxVal int) []int {
	spf := make([]int, maxVal+1)

	for i := 0; i <= maxVal; i++ {
		spf[i] = i
	}

	for i := 2; i*i <= maxVal; i++ {
		if spf[i] == i {
			for j := i * i; j <= maxVal; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
	return spf
}

func getUniquePrimeFactors(x int, spf []int) []int {
	factors := []int{}

	for x > 1 {
		p := spf[x]
		factors = append(factors, p)

		for x%p == 0 {
			x /= p
		}
	}

	return factors
}