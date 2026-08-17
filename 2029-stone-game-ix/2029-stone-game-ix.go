func stoneGameIX(stones []int) bool {
    if len(stones) == 1 {
        return false
    }

    cnt0 := 0
    cnt1 := 0
    cnt2 := 0

    for _, s := range stones {
        switch s % 3 {
            case 0: cnt0++
            case 1: cnt1++
            case 2: cnt2++
        }
    }

    if cnt0 % 2 == 0 {
        return cnt1 > 0 && cnt2 > 0
    }

    
    return abs(cnt1-cnt2) > 2
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
