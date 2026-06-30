func numberOfSubstrings(s string) int {
    n := len(s)

    aCnt, bCnt, cCnt := 0, 0, 0
    res := 0

    for l, r := 0, 0; r < n; r++ {
        ch := s[r]

        switch ch {
        case 'a': aCnt++
        case 'b': bCnt++
        case 'c': cCnt++
        }

        for aCnt > 0 && bCnt > 0 && cCnt > 0 {
            res += n - r
            
            ch = s[l]
            switch ch {
            case 'a': aCnt--
            case 'b': bCnt--
            case 'c': cCnt--
            }

            l++
        }
    }

    return res
}