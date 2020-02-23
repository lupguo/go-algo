package list

import (
	"fmt"
	"math/rand"
)

//ListNode 链表，包含节点元素以及指向下个节点的指针
type ListNode struct {
	Val  int
	Next *ListNode
}

//Generate 随机生成一个指定长度的链表，并返回链表头节点
func Generate(len int) (head *ListNode) {
	head = &ListNode{
		Next: nil,
	}
	for i, p := 0, head ; i < len; i++ {
		p.Next = &ListNode{
			Val:  rand.Intn(10),
			Next: nil,
		}
		p = p.Next
	}
	return
}

//GenerateSorted 生成有序链表数据
func GenerateSorted(len int) (head *ListNode) {
	head = &ListNode{
		Next: nil,
	}
	weight := rand.Intn(5)
	for i, p := 0, head; i < len; i++ {
		p.Next = &ListNode{
			Val:  i+weight,
			Next: nil,
		}
		p = p.Next
	}
	return
}

//Print 打印链表
func (n *ListNode) String() (s string) {
	for p := n; p.Next != nil; p = p.Next {
		s += fmt.Sprintf("%d->", p.Next.Val)
	}
	s += "nil"
	return
}
