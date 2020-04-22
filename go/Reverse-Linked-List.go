/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var p1 *ListNode
    p2 := head
    for p2 != nil {
        temp := p2.Next
        p2.Next = p1
        p1 = p2
        p2 = temp
    }
    return p1
}