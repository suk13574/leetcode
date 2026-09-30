func maxDepth(s string) int {
    res := 0
    depth := 0

    for _, ch := range s {
        if ch == '(' {
            depth++
            res = max(res, depth)
        } else if ch == ')' {
            depth--
        }
    }

    return res
}