package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }

    reverted := 0
    for x > reverted {
        reverted = reverted*10 + x%10
        x /= 10
    }

    return x == reverted || x == reverted/10
}

func main() {
    fmt.Println(isPalindrome(121))  // Output: true
    fmt.Println(isPalindrome(-121)) // Output: false
    fmt.Println(isPalindrome(10))   // Output: false
    fmt.Println(isPalindrome(0))    // Output: true
}
