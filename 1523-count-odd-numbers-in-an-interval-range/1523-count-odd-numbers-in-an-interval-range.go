func countOdds(low int, high int) int {
    cnt := high - low + 1
    if low%2 == 1 && high%2 == 1 {
        return cnt/2 + 1
    }
    return cnt/2
}