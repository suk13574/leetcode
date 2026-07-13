import "strconv"

func sequentialDigits(low int, high int) []int {
    sequential := "123456789"

    minDigit := 0
    maxDigit := 0

    x := low
    for x > 0 {
        x /= 10
        minDigit++
    }

    x = high
    for x > 0 {
        x /= 10
        maxDigit++
    }

    res := []int{}
    for length := minDigit; length <= maxDigit; length++ {
        for l := 0; l+length <= len(sequential); l++ {
            r := l + length
            s := sequential[l:r]
            num, _ := strconv.Atoi(s)

            if low <= num && num <= high {
                res = append(res, num)
            }
        }
    }

    return res
}