func doesAliceWin(s string) bool {
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case 'a', 'e', 'i', 'o', 'u':
            return true
        }
    }
    return false
}
