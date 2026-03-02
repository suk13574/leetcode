func minPartitions(n string) int {
    maxNum := 0
    for i := 0; i < len(n); i++ {
        num := n[i] - '0'
        maxNum = max(maxNum, int(num))
    }

    return maxNum
}