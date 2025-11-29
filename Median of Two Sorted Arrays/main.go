package main

import (
    "fmt"
    "sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Step 1: merge two arrays
    merged := append(nums1, nums2...)
    sort.Ints(merged)

    n := len(merged)
    if n%2 == 1 {
        // odd length, median is middle element
        return float64(merged[n/2])
    } else {
        // even length, median is average of two middle elements
        return float64(merged[n/2-1]+merged[n/2]) / 2.0
    }
}

func main() {
    nums1 := []int{1, 3}
    nums2 := []int{2}
    fmt.Println("Median:", findMedianSortedArrays(nums1, nums2)) // Output: 2.0

    nums3 := []int{1, 2}
    nums4 := []int{3, 4}
    fmt.Println("Median:", findMedianSortedArrays(nums3, nums4)) // Output: 2.5
}
