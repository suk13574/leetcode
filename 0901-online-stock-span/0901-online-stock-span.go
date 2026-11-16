type StockSpanner struct {
    stock []int
}


func Constructor() StockSpanner {
    return StockSpanner{
        stock: make([]int, 0),
    }
}


func (s *StockSpanner) Next(price int) int {
    res := 1

    for pos := len(s.stock) - 1; pos >= 0; pos-- {
        if s.stock[pos] <= price {
            res++
        } else {
            break
        }
    }

    s.stock = append(s.stock, price)

    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */