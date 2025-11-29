package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }

    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }

    return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    interval := 1
    n := len(lists)

    for interval < n {
        for i := 0; i+interval < n; i += interval * 2 {
            lists[i] = mergeTwoLists(lists[i], lists[i+interval])
        }
        interval *= 2
    }

    return lists[0]
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
    l1 := createList([]int{1,4,5})
    l2 := createList([]int{1,3,4})
    l3 := createList([]int{2,6})

    lists := []*ListNode{l1, l2, l3}
    merged := mergeKLists(lists)
    printList(merged)
}
