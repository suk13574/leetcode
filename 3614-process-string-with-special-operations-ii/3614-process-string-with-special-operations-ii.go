func processStr(s string, k int64) byte {
    n := len(s)
    lens := make([]int64, n)

    var length int64

    for i := 0; i < n; i++ {
        ch := s[i]

        switch ch {
        case '*':
            length = max(length-1, 0)
        case '#':
            length *= 2
        case '%':
        default:
            length++
        }

        lens[i] = length
    }

    if k >= length {
        return '.'
    }

    for i := n-1; i >= 0; i-- {
        ch := byte(s[i])

        prevLen := int64(0)
        if i > 0 {
            prevLen = lens[i-1]
        }

        switch ch {
        case '*':

        case '#':
            k %= prevLen
        case '%':
            k = prevLen - 1 - k
        default:
            if k == prevLen {
                return ch
            }
        }
    }

    return '.'
}