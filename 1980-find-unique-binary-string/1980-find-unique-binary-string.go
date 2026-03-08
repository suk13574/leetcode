func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    hi := 1 << n

    makeInt := func(num string) int {
        x := 0
        for _, ch := range num {
            x <<= 1
            if ch == '1' {
                x |= 1
            }
        }
        return x
    }

    makeStr := func(num int) string {
        res := make([]byte, n)
        for i := n-1; i >= 0; i-- {
            if num &1 == 1 {
                res[i] = '1'
            } else {
                res[i] = '0'
            }
            num >>= 1
        }
        return string(res)
    }

    has := make([]bool, hi)
    for _, num := range nums {
        has[makeInt(num)] = true
    }

    for i := 0; i < hi; i++ {
        if !has[i] {
            return makeStr(i)
        }
    }
    return ""
}