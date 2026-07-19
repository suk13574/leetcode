func smallestSubsequence(s string) string {
    n := len(s)

    lastIdx := make([]int, 26)
    for i := 0; i < n; i++ {
        b := s[i] - 'a'
        lastIdx[b] = i
    }

    stack := make([]byte, 0, 26)
    visited := make([]bool, 26)

    for i := 0; i < n; i++ {
        c := s[i]
        b := c - 'a'

        if visited[b] {
            continue
        }

        for len(stack) > 0 {
            top := stack[len(stack)-1]

            if top <= c {
                break
            }

            if lastIdx[top-'a'] <= i {
                break
            }

            stack = stack[:len(stack)-1]
            visited[top-'a'] = false
        }

        stack = append(stack, c)
        visited[b] = true
    }

    return string(stack)
}