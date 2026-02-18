func hasAlternatingBits(n int) bool {
    prev := n & 1
    n >>= 1

    for n > 0 {
        now := n & 1
        if prev == now {
            return false
        }

        prev = now
        n >>= 1
    }

    return true
}