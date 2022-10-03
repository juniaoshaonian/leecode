package Reverse_String_II_541

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n := len(bytes)
	for i := 0; i < n; i += 2 * k {
		if n-i-1 < 2*k {
			reverse(bytes[i:], k)
			break
		} else {
			reverse(bytes[i:i+2*k], k)
		}
	}
	return string(bytes)
}

func reverse(b []byte, k int) {
	n := len(b)
	if n < k {
		k = n
	}
	for i := 0; i < k/2; i++ {
		b[i], b[k-i-1] = b[k-i-1], b[i]
	}
}
