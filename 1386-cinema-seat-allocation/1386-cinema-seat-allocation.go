func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := make(map[int]int)

	for _, r := range reservedSeats {
		row, seat := r[0], r[1]

		if seat >= 2 && seat <= 9 {
			rows[row] |= 1 << seat
		}
	}

	leftMask := 0b111100      // 2,3,4,5
	middleMask := 0b11110000  // 4,5,6,7
	rightMask := 0b1111000000 // 6,7,8,9

	res := n * 2
	for _, mask := range rows {
		left := mask&leftMask == 0
		middle := mask&middleMask == 0
		right := mask&rightMask == 0

		if left && right {
			continue
		}

		if left || middle || right {
			res--
		} else {
			res -= 2
		}
	}

	return res
}