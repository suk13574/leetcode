func twoEditWords(queries []string, dictionary []string) []string {
    n := len(queries[0])
    res := []string{}
    for _, q := range queries {
        for _, d := range dictionary {
            diff := 0

            for i := 0; i < n; i++ {
                if q[i] != d[i] {
                    diff++
                }

                if diff > 2 {
                    break
                }
            }
            
            if diff <= 2 {
                res = append(res, q)
                break
            }
        }
    }

    return res
}