func shortestBeautifulSubstring(s string, k int) string {
    var res string
    ones := 0
    
    l := 0
    for r := 0; r < len(s); r++ {
        if s[r] == '1' {
            ones++
        }

        for ones == k {
            sub := s[l:r+1]

            if res == "" || len(sub) < len(res) || len(sub) == len(res) && sub < res  {
                res = sub
            }

            if s[l] == '1' {
                ones--
            }

            l++
        }
    }
    
    return res
}