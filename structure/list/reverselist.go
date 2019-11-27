package list

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
func ReverseList(head *ListNode) *ListNode {
	//return IterateReverseList(head)
	r, _ := RecursiveReverseList(head.Next)
	return r
}

//IterateReverseList 迭代输出
// 	思路：申请一个反转的头节点r，从标准节点开始检测：r, p->p.Next, r为已逆序排列好的链表，等于逆序链表r和链表p，分四步进行：
//		1. 存储p.Next到tmp防止指针丢失
//		2. 将p.Next执行r，完成p->r的逆序链表链接
//		3. 迭代推进r指向p
//		4. 迭代推进p指向tmp
//	Tips：
//		1. 注意区分空链表、单一节点链表、多节点链表情况；
//		2. 返回时候，注意补充完善head哨兵节点内容；
func IterateReverseList(head *ListNode) (r *ListNode) {
	for p := head.Next; p != nil; {
		tmp := p.Next
		p.Next = r
		r = p
		p = tmp
	}
	return &ListNode{
		Next: r,
	}
}

//RecursiveReverseList 递归反转链表
//	思路: 从head.Next节点开始，按{当前节点p, 子问题p.Next}思路分析
//	Tips: 注意区分 head->nil, head->node(1,nil), head->node(1,x)->node(2,x)...三类情况
//			同时在递归内部，同一使用标准节点，而不掺杂head头哨兵节点
func RecursiveReverseList(p *ListNode) (r *ListNode, tail *ListNode) {
	// 1. 递归退出条件
	if p == nil || p.Next == nil {
		return &ListNode{
			Val:  0,
			Next: p,
		}, p
	}
	// 2. p.Next不为空，求p.Next的反序列列表
	r, tail = RecursiveReverseList(p.Next)
	// 3. 节点指针反转操作
	p.Next = nil
	tail.Next = p
	tail = p
	return
}
