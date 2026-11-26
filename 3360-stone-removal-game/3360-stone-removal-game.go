func canAliceWin(n int) bool {
    for i := 10; i > 0; i-- {
        if n < i {
            return i % 2 == 1
        }
        n -= i
    }

    return false
}