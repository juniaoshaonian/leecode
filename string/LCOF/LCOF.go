package LCOF

func reverseLeftWords(s string, n int) string {
	s = s + s
	k := len(s)
	n = n % k
	return s[k : k+n]
}
