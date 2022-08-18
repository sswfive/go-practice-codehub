package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"一二三一", 3},
		{"这里是江苏南京南", 7},
		{"黑化肥挥发发挥会花飞挥发发挥灰", 5},
	}
	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s;"+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发挥会花飞挥发发挥灰"
	ans := 5
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+"expected %d", actual, s, ans)
		}
	}
}
