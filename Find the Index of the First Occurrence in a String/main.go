package main

import (
    "fmt"
    "strings"
)

func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}

func main() {
    fmt.Println(strStr("sadbutsad", "sad"))    // 0
    fmt.Println(strStr("leetcode", "leeto"))   // -1
}
