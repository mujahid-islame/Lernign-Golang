package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    k := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}

func main() {
    nums1 := []int{1,1,2}
    k1 := removeDuplicates(nums1)
    fmt.Println(k1, nums1[:k1]) // 2 [1 2]

    nums2 := []int{0,0,1,1,1,2,2,3,3,4}
    k2 := removeDuplicates(nums2)
    fmt.Println(k2, nums2[:k2]) // 5 [0 1 2 3 4]
}
