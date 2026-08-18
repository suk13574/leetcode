func largestInteger(nums []int, k int) int {
    n := len(nums)

    maxNum := 0
    numCnt := make(map[int]int)

    for _, num := range nums {
        maxNum = max(maxNum, num)
        numCnt[num]++
    }

    if k == n {
        return maxNum
    }

    if k == 1 {
        res := -1
        for _, num := range nums {
            if numCnt[num] == 1 {
                res = max(res, num)
            }
        }

        return res
    }

    left := nums[0]
    right := nums[n-1]

    if numCnt[left] >= 2 && numCnt[right] >= 2 {
        return -1
    }

    if numCnt[left] >= 2 {
        return right
    } else if numCnt[right] >= 2 {
        return left
    }

    return max(left, right)
}