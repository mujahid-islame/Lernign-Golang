package main

import "fmt"

func removeElement(nums []int, val int) int {
    k := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}

func main() {
    nums1 := []int{3,2,2,3}
    k1 := removeElement(nums1, 3)
    fmt.Println(k1, nums1[:k1]) // 2 [2 2]

    nums2 := []int{0,1,2,2,3,0,4,2}
    k2 := removeElement(nums2, 2)
    fmt.Println(k2, nums2[:k2]) // 5 [0 1 3 0 4] (order may vary)
}
