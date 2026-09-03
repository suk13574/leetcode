func uniformArray(nums1 []int) bool {
    const INF int = math.MaxInt
    
    minOdd := INF
    minEven := INF

    for _, num := range nums1 {
        if num % 2 == 0 {
            minEven = min(minEven, num)
        } else {
            minOdd = min(minOdd, num)
        }
    }

    if minOdd != INF && minEven != INF && minEven < minOdd {
        return false
    }

    return true
}