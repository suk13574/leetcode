func maxSumOfSquares(num int, sum int) string {
    n := make([]byte, num)

    for i := 0; i < num; i++ {
        if sum > 9 {
            n[i] = '9'
            sum -= 9
        } else {
            n[i] = '0' + byte(sum)
            sum = 0
        }
    }

    if sum != 0 {
        return ""
    }

    return string(n)
}