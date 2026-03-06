func checkOnesSegment(s string) bool {
    zeroIdx := len(s)

    for i := 1; i < len(s); i++ {
        if s[i] == '0' {
            zeroIdx = i
            break
        }
    }

    for i := zeroIdx+1; i < len(s); i++ {
        if s[i] == '1' {
            return false
        }
    }

    return true
}