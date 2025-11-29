package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    } else if list2 != nil {
        current.Next = list2
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
    l1 := createList([]int{1,2,4})
    l2 := createList([]int{1,3,4})
    merged := mergeTwoLists(l1, l2)
    printList(merged)

    l3 := createList([]int{})
    l4 := createList([]int{})
    printList(mergeTwoLists(l3, l4))

    l5 := createList([]int{})
    l6 := createList([]int{0})
    printList(mergeTwoLists(l5, l6))
}
