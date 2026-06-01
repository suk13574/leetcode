func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

    res := 0
    sell := 0
    for i := 0; i < len(cost); i++ {
        if sell >= 2 {
            sell = 0
            continue
        }

        sell++
        res += cost[i]
    }
    
    return res
}