import (
	"fmt"
)

func maxFreqSum(s string) int {
	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	alphabet := make([]int, 26)

	vMax := 0
	cMax := 0
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		alphabet[idx]++

		if isVowel(s[i]) {
			vMax = max(vMax, alphabet[idx])
		} else {
			cMax = max(cMax, alphabet[idx])
		}
	}

	return vMax + cMax
}