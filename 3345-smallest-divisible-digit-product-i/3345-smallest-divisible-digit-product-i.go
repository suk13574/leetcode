func smallestNumber(n int, t int) int {
	for i := 0; i < 10; i++ {
        p := productDigits(n+i)
        if p % t == 0 {
            return n+i
        }
    }

    return -1
}

func productDigits(n int) int {
    res := 1
    for n > 0 {
        res *= n%10
        n /= 10
    }

    return res
}