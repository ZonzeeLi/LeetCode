package main

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(t) && j < len(s) {
		if s[j] == t[i] {
			j++
			if j == len(s) {
				return true
			}
		}
		i++
	}
	return j == len(s)
}
