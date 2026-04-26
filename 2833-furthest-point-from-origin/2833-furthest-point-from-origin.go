func furthestDistanceFromOrigin(moves string) int {
    res := 0
    tmp := 0
    for i := 0; i < len(moves); i++ {
        d := moves[i]

        if d == 'L' {
            res--
        } else if d == 'R' {
            res++
        } else {
            tmp++
        }
    }

    res = abs(res)
    return res + tmp
}

func abs(x int) int {
    if x < 0 {
        x = -x
    }
    return x
}