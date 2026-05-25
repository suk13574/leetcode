func canReach(s string, minJump int, maxJump int) bool {
    n := len(s)
    q := []int{0}
    farhest := 0 

    for cur := 0; cur < len(q); cur++ {
        i := q[cur]
        
        if i == n-1 {
            return true
        }

        start := max(i+minJump, farhest+1)
        end := min(i+maxJump, n-1)
        for j := start; j <= end; j++ {
            if s[j] == '1' {
                continue
            }
            q = append(q, j)
        }

        farhest = max(farhest, end)
    }

    return false
}