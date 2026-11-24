package main

import "fmt"

// IsMatch reports whether s matches pattern p entirely, where '.' matches
// any single character and '*' means zero or more of the preceding element.
// Classic table DP over (prefix of s, prefix of p).
func IsMatch(s, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	// Empty s vs pattern: only chains like a*b*c* can vanish.
	for j := 2; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	single := func(sc, pc byte) bool { return pc == '.' || pc == sc }
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || // zero occurrences of p[j-2]
					(single(s[i-1], p[j-2]) && dp[i-1][j]) // one more occurrence
			} else {
				dp[i][j] = single(s[i-1], p[j-1]) && dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func main() {
	fmt.Println(IsMatch("aa", "a*"), IsMatch("ab", ".*"), IsMatch("aa", "a"))
}
