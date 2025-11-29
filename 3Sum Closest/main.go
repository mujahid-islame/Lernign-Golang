package main

import (
    "fmt"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]

    for i := 0; i < len(nums)-2; i++ {
        left, right := i+1, len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if abs(sum-target) < abs(closest-target) {
                closest = sum
            }

            if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    return closest
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2
    fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))       // 0
}
