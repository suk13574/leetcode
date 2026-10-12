func maxSum(nums []int, k int) int {
    var MOD int64 = 1_000_000_007

    square := func(x int) int64 {
        return (int64(x)*int64(x))%MOD
    }

    count := [30]int{}
    
    for _, num := range nums {
        x := num
        pos := 0
        for x > 0 {
            if x & 1 == 1 {
                count[pos]++
            }

            x >>= 1
            pos++
        }
    }

    res := int64(0)
    for ;k > 0; k-- {
        x := 0
        
        for i := 0; i < 30; i++ {
            if count[i] == 0 {
                continue
            }

            x |= (1 << i)
            count[i]--
        }

        res += square(x)
        res %= MOD
    }

    return int(res%MOD)
}