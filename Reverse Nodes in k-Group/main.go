package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{0, head}
    prevGroupEnd := dummy

    for {
        kth := prevGroupEnd
        for i := 0; i < k && kth != nil; i++ {
            kth = kth.Next
        }
        if kth == nil {
            break
        }

        groupStart := prevGroupEnd.Next
        nextGroupStart := kth.Next

        // Reverse the group
        prev, curr := nextGroupStart, groupStart
        for curr != nextGroupStart {
            tmp := curr.Next
            curr.Next = prev
            prev = curr
            curr = tmp
        }

        prevGroupEnd.Next = kth
        prevGroupEnd = groupStart
    }

    return dummy.Next
}

// Helper to create list
func createList(nums []int) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for _, v := range nums {
        current.Next = &ListNode{Val: v}
        current = current.Next
    }
    return dummy.Next
}

// Helper to print list
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " ")
        head = head.Next
    }
    fmt.Println()
}

func main() {
    l1 := createList([]int{1,2,3,4,5})
    printList(reverseKGroup(l1, 2)) // Output: 2 1 4 3 5

    l2 := createList([]int{1,2,3,4,5})
    printList(reverseKGroup(l2, 3)) // Output: 3 2 1 4 5
}
