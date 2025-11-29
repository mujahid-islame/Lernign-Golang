package main

import "fmt"

func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true

    // Initialize dp for patterns like a*, a*b*, a*b*c* etc
    for j := 2; j <= n; j += 2 {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-2]
        }
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '.' || p[j-1] == s[i-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-2] || ((p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j])
            }
        }
    }
    return dp[m][n]
}

func main() {
    fmt.Println(isMatch("aa", "a"))   // Output: false
    fmt.Println(isMatch("aa", "a*"))  // Output: true
    fmt.Println(isMatch("ab", ".*"))  // Output: true
    fmt.Println(isMatch("mississippi", "mis*is*p*.")) // Output: false
}
