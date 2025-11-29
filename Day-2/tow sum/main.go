package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // মান এবং index রাখার জন্য map
    numMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; ok {
            // যদি complement map-এ থাকে, তাহলে আমরা সমাধান পেয়েছি
            return []int{j, i}
        }
        // map-এ current num রাখো
        numMap[num] = i
    }
    return nil // কোনো সমাধান না থাকলে (theoretically এখানে আসবে না)
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) // Output: [0 1]
}
