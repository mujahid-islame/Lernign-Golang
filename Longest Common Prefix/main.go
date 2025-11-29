package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0]

    for _, word := range strs {
        for len(prefix) > 0 && len(word) < len(prefix) || word[:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1]
        }
        if prefix == "" {
            return ""
        }
    }
    return prefix
}

func main() {
    fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
    fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
