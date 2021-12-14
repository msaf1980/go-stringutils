package stringutils

// Reverse returns reversed string (rune-wise left to right).
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
