func maxDepthAfterSplit(seq string) []int {
    res := []int{}
    depth := 0

    for _, ch := range seq {
        if ch == '(' {
            depth++
        }

        if depth%2 == 1 {
            res = append(res, 0)
        } else {
            res = append(res, 1)
        }


        if ch == ')' {
            depth--
        }
    }


    return res
}