package main

import "fmt"

// Linked list node
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to add two numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{} // dummy node to simplify code
    current := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    return dummy.Next
}

// Helper function to create linked list from slice
func createList(nums []int) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for _, num := range nums {
        current.Next = &ListNode{Val: num}
        current = current.Next
    }
    return dummy.Next
}

// Helper function to print linked list
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val)
        if head.Next != nil {
            fmt.Print(" -> ")
        }
        head = head.Next
    }
    fmt.Println()
}

func main() {
    l1 := createList([]int{2, 4, 3})
    l2 := createList([]int{5, 6, 4})

    result := addTwoNumbers(l1, l2)
    printList(result) // Output: 7 -> 0 -> 8
}
