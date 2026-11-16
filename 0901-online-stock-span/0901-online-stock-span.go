type Stock struct {
	price int
	span  int
}

type StockSpanner struct {
	stock []*Stock
}

func Constructor() StockSpanner {
	return StockSpanner{
		stock: make([]*Stock, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	span := 1

	for len(s.stock) > 0 {
		n := len(s.stock)

		top := s.stock[n-1]

		if top.price <= price {
			span += top.span
			s.stock = s.stock[:n-1]
		} else {
			break
		}
	}

	s.stock = append(s.stock, &Stock{
		price: price,
		span:  span,
	})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */