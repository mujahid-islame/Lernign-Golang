package main

import "fmt"

func reverse(x int) int {
    var result int
    for x != 0 {
        pop := x % 10
        x /= 10

        if result > 214748364 || (result == 214748364 && pop > 7) {
            return 0
        }
        if result < -214748364 || (result == -214748364 && pop < -8) {
            return 0
        }

        result = result*10 + pop
    }
    return result
}

func main() {
    fmt.Println(reverse(123))   // Output: 321
    fmt.Println(reverse(-123))  // Output: -321
    fmt.Println(reverse(120))   // Output: 21
    fmt.Println(reverse(1534236469)) // Output: 0 (overflow)
}
