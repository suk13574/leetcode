func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	needBoth := min(need1, need2)

	var res int64 = 0

	res += int64(needBoth) * int64(min(costBoth, cost1+cost2))
	res += int64(need1-needBoth) * int64(min(costBoth, cost1))
	res += int64(need2-needBoth) * int64(min(costBoth, cost2))

	return res
}