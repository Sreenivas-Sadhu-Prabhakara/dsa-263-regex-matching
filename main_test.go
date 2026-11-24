package main

import "testing"

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true}, // c* vanishes, a* eats two
		{"mississippi", "mis*is*p*.", false},
		{"mississippi", "mis*is*ip*.", true},
		{"", "", true},
		{"", "a*b*", true},
		{"", ".", false},
		{"abc", "abc", true},
		{"aaa", "a*a", true}, // star must leave one a for the literal
		{"ab", ".*c", false},
	}
	for _, c := range cases {
		if got := IsMatch(c.s, c.p); got != c.want {
			t.Errorf("IsMatch(%q,%q) = %v, want %v", c.s, c.p, got, c.want)
		}
	}
}
