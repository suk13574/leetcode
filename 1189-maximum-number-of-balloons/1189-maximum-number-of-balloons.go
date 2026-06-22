func maxNumberOfBalloons(text string) int {
    byteToVal := map[byte]int {
        'b': 0,
        'a': 1,
        'l': 2,
        'o': 3,
        'n': 4,
    }

    ballons := make([]int, 5)
    for i := 0; i < len(text); i++ {
        b := text[i]

        if v, ok := byteToVal[b]; ok {
            ballons[v]++
        }
    }

    res := 1 << 32
    for _, b := range []byte{'b', 'a', 'l', 'o', 'n'} {
        cnt := ballons[byteToVal[b]]
        if b == 'l' || b == 'o' {
            cnt /= 2
        }

        res = min(res, cnt)
    }

    return res
}