package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    slow, fast := dummy, dummy

    // Move fast n+1 steps ahead
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Move both until fast reaches the end
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    // Remove the nth node
    slow.Next = slow.Next.Next

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
    head1 := createList([]int{1,2,3,4,5})
    result1 := removeNthFromEnd(head1, 2)
    printList(result1)

    head2 := createList([]int{1})
    result2 := removeNthFromEnd(head2, 1)
    printList(result2)

    head3 := createList([]int{1,2})
    result3 := removeNthFromEnd(head3, 1)
    printList(result3)
}
