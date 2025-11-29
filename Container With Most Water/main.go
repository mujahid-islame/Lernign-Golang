package main

import "fmt"

func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxWater := 0

    for left < right {
        h := height[left]
        if height[right] < h {
            h = height[right]
        }

        area := h * (right - left)
        if area > maxWater {
            maxWater = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}

func main() {
    fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7})) // Output: 49
    fmt.Println(maxArea([]int{1,1}))               // Output: 1
}
