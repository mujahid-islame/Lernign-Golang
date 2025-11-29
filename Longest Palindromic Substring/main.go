package main

import "fmt"

// Function to expand around center and return palindrome
func expandAroundCenter(s string, left, right int) string {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left+1 : right]
}

func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }

    longest := ""

    for i := 0; i < len(s); i++ {
        // odd length palindrome
        odd := expandAroundCenter(s, i, i)
        if len(odd) > len(longest) {
            longest = odd
        }

        // even length palindrome
        even := expandAroundCenter(s, i, i+1)
        if len(even) > len(longest) {
            longest = even
        }
    }

    return longest
}

func main() {
    testCases := []string{"babad", "cbbd", "a", "ac"}

    for _, s := range testCases {
        fmt.Printf("Input: %s, Longest Palindromic Substring: %s\n", s, longestPalindrome(s))
    }
}
