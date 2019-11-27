package list

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	FindAndRemoveNode(head,  n)
	return head
}

//FindAndRemoveNode 找到并移除节点
//	思路: 利用递归原理，从最后一个节点开始level开始统计，每返回一层，level+1；
//		当level+1=n时候，node.Next即为需要被删除的元素
func FindAndRemoveNode(p *ListNode, n int) (level int) {
	// 1. 倒数第1个元素，即p.Next为空的元素
	if p.Next == nil {
		return 1
	}
	// 2. 非倒数第一个元素
	level = FindAndRemoveNode(p.Next, n)
	// 3. 迭代返回时候，层数+1
	level++
	// 4. 当level==n+1时，p.next即需要被移除的元素
	if level == n+1 {
		p.Next = p.Next.Next
	}
	return level
}
