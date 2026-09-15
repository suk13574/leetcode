func maxPalindromes(s string, k int) int {
    res := 0

    check := func(l, r int) bool {
        for l < r {
            if s[l] != s[r] {
                return false
            }

            l++
            r--
        }
        return true
    }

    for i := 0; i+k <= len(s); {
        // length k        
        if check(i, i+k-1) {
            res++
            i = i+k
            continue
        }

        // length k+1
        if i+k < len(s) && check(i, i+k) {
            res++
            i = i+k+1
            continue
        }

        i++
    }

    return res
}