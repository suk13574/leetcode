func countPermutations(complexity []int) int {
    const MOD = 1000000007

    minValue := complexity[0]
    var res int64 = 1
    for i := 1; i < len(complexity); i++ {
        v := complexity[i]
        if v <= minValue {
            return 0
        }

        res *= int64(i)
        res %= MOD

        minValue = min(minValue, v)
    }

    return int(res)
}