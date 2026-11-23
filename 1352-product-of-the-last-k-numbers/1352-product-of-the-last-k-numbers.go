type ProductOfNumbers struct {
	prefix []int
}

func Constructor() ProductOfNumbers {
	p := []int{1}
	return ProductOfNumbers{prefix: p}
}

func (pn *ProductOfNumbers) Add(num int) {
	if num == 0 {
		pn.prefix = []int{1}
	} else {
		n := len(pn.prefix)
		pn.prefix = append(pn.prefix, pn.prefix[n-1]*num)
	}
}

func (pn *ProductOfNumbers) GetProduct(k int) int {
	n := len(pn.prefix)

	if n <= k {
		return 0
	}

	return pn.prefix[n-1] / pn.prefix[n-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */