func longestCommonPrefix(arr1 []int, arr2 []int) int {
    prefix := make(map[int]struct{})

    for _, v := range arr1 {
        for v > 0 {
            prefix[v] = struct{}{}
            v /= 10
        }
    }

    numLen := func(x int) int {
        l := 0

        for x > 0 {
            x /= 10
            l++
        }
        return l
    }

    res := 0
    for _, v := range arr2 {
        l := numLen(v)
        for v > 0 {
            if _, ok := prefix[v]; ok {
                res = max(res, l)
                break
            }
            l--
            v /= 10
        }
    }

    return res
}