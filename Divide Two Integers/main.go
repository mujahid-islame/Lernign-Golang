package main

import (
    "fmt"
    "math"
)

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    negative := (dividend < 0) != (divisor < 0)

    dvd := abs(dividend)
    dvs := abs(divisor)
    result := 0

    for dvd >= dvs {
        temp, multiple := dvs, 1
        for dvd >= temp<<1 {
            temp <<= 1
            multiple <<= 1
        }
        dvd -= temp
        result += multiple
    }

    if negative {
        return -result
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    fmt.Println(divide(10, 3))   // 3
    fmt.Println(divide(7, -3))   // -2
    fmt.Println(divide(-2147483648, -1)) // 2147483647
}
