func compareVersion(version1 string, version2 string) int {
	idx1, idx2 := 0, 0
	n, m := len(version1), len(version2)

	for idx1 < n || idx2 < m {
		v1 := 0
		for idx1 < n && version1[idx1] != '.' {
			v1 = (v1 * 10) + int(version1[idx1]-'0')
			idx1++
		}

		v2 := 0
		for idx2 < m && version2[idx2] != '.' {
			v2 = (v2 * 10) + int(version2[idx2]-'0')
			idx2++
		}

		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}

		if idx1 < n && version1[idx1] == '.' {
			idx1++
		}
		if idx2 < m && version2[idx2] == '.' {
			idx2++
		}
	}

	return 0
}
