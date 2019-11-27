package list

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
//
// Related Topics 链表

//MergeTwoLists 合并两个有序列表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	m := new(ListNode)
	head := m
	p, q := l1.Next, l2.Next
	for p != nil && q != nil {
		if p.Val <= q.Val {
			m.Next = p
			p = p.Next
		} else {
			m.Next = q
			q = q.Next
		}
		m = m.Next
	}
	// 剩余列表的合并
	if q == nil {
		m.Next = p
	} else {
		m.Next = q
	}
	return head
}
