import (
	"math"
)

func canCommuicate(u, v int, n int, langMap [][]bool) bool {
	for l := 1; l <= n; l++ {
		if langMap[u][l] && langMap[v][l] {
			return true
		}
	}
	return false
}

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	langMap := make([][]bool, m+1)
	for i := 0; i < len(langMap); i++ {
		langMap[i] = make([]bool, n+1)
	}

	for i, langs := range languages {
		for _, l := range langs {
			langMap[i+1][l] = true
		}
	}

	teachFriend := map[int]bool{}
	for _, f := range friendships {
		u, v := f[0], f[1]
		if !canCommuicate(u, v, n, langMap) {
			teachFriend[u] = true
			teachFriend[v] = true
		}
	}

	answer := math.MaxInt
	for l := 1; l <= n; l++ {
		cnt := 0
		for u := range teachFriend {
			if !langMap[u][l] {
				cnt++
			}
		}
		answer = min(answer, cnt)
	}

	return answer
}