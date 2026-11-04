    func minSwapsCouples(row []int) int {
        n := len(row)

        res := 0

        for i := 1; i < n; i += 2 {
            person := row[i-1]
            cur := row[i]

            if (person ^ cur) == 1 {
                continue
            }
            
            res++

            for j := i+1; j < n; j++ {
                if (person ^ row[j]) == 1 {
                    row[i], row[j] = row[j], row[i]
                    break
                }
            }
        }

        return res
    }