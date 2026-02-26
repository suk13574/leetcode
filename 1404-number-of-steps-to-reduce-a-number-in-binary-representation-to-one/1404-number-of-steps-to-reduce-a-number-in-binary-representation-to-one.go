func numSteps(s string) int {
    n := 0
    for i := 0; i < len(s); i++ {
        n = (n << 1) | int(s[i]-'0')
    }

    res := 0
    for n > 1 {
        if n & 1 == 1 {
            res++
            n += 1
        }
        n >>= 1
        res++
    }

    return res
}