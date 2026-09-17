func minSumOfLengths(arr []int, target int) int {
    const INF = 1 << 30

    n := len(arr)

    best := make([]int, n)
    for i := 0; i < n; i++ {
        best[i] = INF
    }

    res := INF
    sum := 0
    l := 0
    for r := 0; r < n; r++ {
        sum += arr[r]

        for sum > target {
            sum -= arr[l]
            l++
        }

        if r > 0 {
            best[r] = best[r-1]
        }

        if sum == target {
            length := r - l + 1

            best[r] = min(best[r], length)

            if l > 0 && best[l-1] != INF {
                res = min(res, length+best[l-1])
            }
        }
    }

    if res == INF {
        return -1
    }

    return res
}