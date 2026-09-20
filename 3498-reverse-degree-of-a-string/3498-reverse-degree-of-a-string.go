func reverseDegree(s string) int {
    res := 0
    for i, ch := range s {
        pos := int(ch - 'a')
        res += (i+1) * (26-pos)
    }

    return res
}