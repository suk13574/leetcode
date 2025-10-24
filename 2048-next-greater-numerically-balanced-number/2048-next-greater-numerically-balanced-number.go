import "sort"


func nextBeautifulNumber(n int) int {
	cand := genBalancedNums()
	sort.Ints(cand)

	for _, num := range cand {
		if num > n {
			return num
		}
	}
	return -1
}

func genCom() [][]int {
	result := [][]int{}
	current := []int{}

	var backtrack func(digit int, totlaDigits int)
	backtrack = func(digit int, totlaDigits int) {
		if totlaDigits > 7 {
			return
		}

		if len(current) > 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
		}

		for d := digit; d <= 6; d++ {
			if totlaDigits <= 7 {
				for i := 0; i < d; i++ {
					current = append(current, d)
				}
				backtrack(d+1, totlaDigits+d)
				current = current[:len(current)-d]
			}
		}
	}

	backtrack(1, 0)

	return result
}

func genBalancedNums() []int {
	result := []int{}

	coms := genCom()

	for _, digits := range coms {
		perms := genPermutations(digits)
		for _, perm := range perms {
			if perm[0] == 0 {
				continue
			}
			num := sliceToNumber(perm)
			if num > 0 && num <= 10000000 {
				result = append(result, num)
			}
		}
	}
	return result
}

func genPermutations(nums []int) [][]int {
	result := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		seen := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			if used[i] || seen[nums[i]] {
				continue
			}

			seen[nums[i]] = true
			used[i] = true
			current = append(current, nums[i])

			backtrack()

			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

func sliceToNumber(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		result *= 10
		result += nums[i]
	}
	return result
}