package stringutils

import (
	"strings"
	"testing"
)

func Test_Split2(t *testing.T) {
	tests := []struct {
		s           string
		sep         string
		want0       string
		want1       string
		wantStrings int
	}{
		{"", "", "", "", 1},
		{"", "&", "", "", 1},
		{"test", "&", "test", "", 1},
		{"test&", "&", "test", "", 2},
		{"test&after", "&", "test", "after", 2},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+tt.sep, func(t *testing.T) {
			s0, s1, n := Split2(tt.s, tt.sep)
			if s0 != tt.want0 {
				t.Errorf("Split2() s[0] = %v, want %v", s0, tt.want0)
			}
			if s1 != tt.want1 {
				t.Errorf("Split2() s[1] = %v, want %v", s1, tt.want1)
			}
			if n != tt.wantStrings {
				t.Errorf("Split2() count = %v, want %v", n, tt.wantStrings)
			}
		})
	}
}

// for compare with Benchmark_Split2
func Benchmark_Split(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strings.SplitN(s, "&", 1)
	}
}

func Benchmark_Split2(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = Split2(s, "&")
	}
}
