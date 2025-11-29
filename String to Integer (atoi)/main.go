package main

import (
    "fmt"
    "math"
    "strings"
)

func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }

    sign := 1
    i := 0
    if s[0] == '+' {
        i++
    } else if s[0] == '-' {
        sign = -1
        i++
    }

    result := 0
    for ; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            break
        }
        digit := int(s[i] - '0')

        if result > (math.MaxInt32-digit)/10 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }

        result = result*10 + digit
    }

    return result * sign
}

func main() {
    fmt.Println(myAtoi("42"))          // Output: 42
    fmt.Println(myAtoi("   -042"))     // Output: -42
    fmt.Println(myAtoi("1337c0d3"))    // Output: 1337
    fmt.Println(myAtoi("0-1"))         // Output: 0
    fmt.Println(myAtoi("words and 987")) // Output: 0
}
