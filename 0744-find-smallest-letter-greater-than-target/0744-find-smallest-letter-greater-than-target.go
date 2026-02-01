func nextGreatestLetter(letters []byte, target byte) byte {
	letter := make([]bool, 26)
	for _, l := range letters {
		letter[l-'a'] = true
	}

	for i := target - 'a' + 1; i < 26; i++ {
		if letter[i] {
			return i + 'a'
		}
	}
	return letters[0]
}