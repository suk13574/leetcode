func binaryGap(n int) int {
    prev := -1
    cur := 0
    res := 0
    for n > 0 {
        now := n & 1
        
        if now == 1 {
            if prev != -1 {
                res = max(res, cur-prev)
            }

            prev = cur
        }
        cur++
        n >>= 1
    }

    return res
}