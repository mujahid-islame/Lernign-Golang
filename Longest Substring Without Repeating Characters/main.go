package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int) // map to store last index of each character
    maxLen := 0
    start := 0 // start index of current substring

    for i, ch := range s {
        if idx, found := charIndex[ch]; found && idx >= start {
            start = idx + 1
        }
        charIndex[ch] = i
        if i-start+1 > maxLen {
            maxLen = i - start + 1
        }
    }

    return maxLen
}

func main() {
    testCases := []string{"abcabcbb", "bbbbb", "pwwkew"}
    
    for _, s := range testCases {
        fmt.Printf("Input: %s, Output: %d\n", s, lengthOfLongestSubstring(s))
    }
}
