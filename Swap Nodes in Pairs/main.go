package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    current := dummy

    for current.Next != nil && current.Next.Next != nil {
        first := current.Next
        second := current.Next.Next

        first.Next = second.Next
        second.Next = first
        current.Next = second

        current = first
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
    l1 := createList([]int{1,2,3,4})
    printList(swapPairs(l1))

    l2 := createList([]int{})
    printList(swapPairs(l2))

    l3 := createList([]int{1})
    printList(swapPairs(l3))

    l4 := createList([]int{1,2,3})
    printList(swapPairs(l4))
}
