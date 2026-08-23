func checkDivisibility(n int) bool {
    sum := 0
    product := 1

    x := n
    for x > 0 {
        v := x % 10
        
        sum += v
        product *= v

        x /= 10
    }

    if n % (sum+product) == 0 {
        return true
    }

    return false
}