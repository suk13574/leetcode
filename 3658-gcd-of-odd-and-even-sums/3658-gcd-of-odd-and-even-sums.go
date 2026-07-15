func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }

    return a
}

func gcdOfOddEvenSums(n int) int {
    oddSum := n*n
    evenSum := (n+1)*n

    return gcd(oddSum, evenSum)   
}