func maxProduct(n int) int {
    digits := []int{}

    x := n
    for x > 0 {
        digits = append(digits, x%10)
        x /= 10
    }

    sort.Sort(sort.Reverse(sort.IntSlice(digits)))

    return digits[0] * digits[1]
}