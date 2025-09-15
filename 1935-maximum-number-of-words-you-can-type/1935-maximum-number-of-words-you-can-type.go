func canBeTypedWords(text string, brokenLetters string) int {
	brokenL := make([]bool, 26)
	for i := 0; i < len(brokenLetters); i++ {
		b := brokenLetters[i]
		brokenL[b-'a'] = true
	}

	text += " "

	answer := 0
	isTyping := true
	for i := 0; i < len(text); i++ {
		s := text[i]
		if s == ' ' {
			if isTyping {
				answer++
			}
			isTyping = true
			continue
		}

		if brokenL[s-'a'] {
			isTyping = false
		}
	}
	return answer
}