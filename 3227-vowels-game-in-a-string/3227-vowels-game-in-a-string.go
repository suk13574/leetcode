func doesAliceWin(s string) bool {
    isVowel := func(c byte) bool {
        switch c {
        case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
            return true
        }
        return false
    }

    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            return  true
        }
    }
    return false
}