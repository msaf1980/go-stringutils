package stringutils

import "strings"

var empthy = ""

// Split2 return the split string results (without memory allocations)
//   If sep string not found: 's' '' 1
//   If s or sep string is empthy: 's' '' 1
//   In other cases: 's0' 's2' 2
func Split2(s string, sep string) (string, string, int) {
	if sep == "" {
		return s, empthy, 1
	}

	if pos := strings.Index(s, sep); pos == -1 {
		return s, empthy, 1
	} else if pos == len(s)-1 {
		return s[0:pos], empthy, 2
	} else {
		return s[0:pos], s[pos+1:], 2
	}
}

func truncate(s *[]string, to int) {
	*s = (*s)[:to]
}

// SplitN return splitted slice (use pre-allocated buffer) and end position (for detect if string contains more fields for split)
func SplitN(s string, sep string, buf []string) ([]string, int) {
	n := len(buf)
	i := 0
	p := 0

	for i < n {
		if pos := strings.Index(s, sep); pos == -1 {
			buf[i] = s
			return buf[0 : i+1], p + len(s)
		} else {
			buf[i] = s[0:pos]
			p += pos + len(sep)
			i++
			if i == n {
				break
			}
			if pos+1 == len(s) {
				buf[i] = s[pos:pos]
				return buf[0 : i+1], p
			}
			s = s[pos+1:]
		}
	}
	return buf[0:n], p
}
