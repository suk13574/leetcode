func sumGame(num string) bool {
    n := len(num)

    firstSum := 0
    secondSum := 0
    firstQ := 0
    secondQ := 0

    for i := 0; i < n; i++ {
        ch := num[i]

        if ch == '?' {
            if i < n/2 {
                firstQ++
            } else {
                secondQ++
            }
            continue
        }

        digit := int(ch-'0')

        if i < n/2 {
            firstSum += digit
        } else {
            secondSum += digit
        }
    }

    if (firstQ+secondQ) % 2 == 1 {
        return true
    }

    return 2*(firstSum-secondSum) != 9*(secondQ-firstQ)
}