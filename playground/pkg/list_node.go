package pkg

type ListNode struct {
	Next *ListNode
	Val  int
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Next: nil,
		Val:  val,
	}
}
