package dstruct

import "fmt"

//PriorityQueue 优先队列
// 术语: 小顶堆, 大顶堆, 元素插入, 移除堆顶元素, 上滤, 下滤
// 实现:
//		链表: 插入O(1)，移除操作, 遍历最小元素O(N)
//		链表: 持续保持排序，O(N)，移除操作，O(1)
//		二叉查找树: 插入,移除 O(LogN)，可能有不平衡的问题，可以通过平衡二叉树解决，但二叉查找树多了很多不必要的操作
//		二叉堆(堆)(2-堆): 使用数组实现，支持插入，移除操作 O(LogN)，线性时间构建N项的优先队列
// 结构性:
// 		1. 完全二叉树(完全填满的二叉树，同时底层的元素从做到右依次填满)，高度为logN
// 		2. 完全二叉树很有规律，用数组存放而不需要指针，数组有一个位置0，i从1开始，任意位置i，其左儿子在位置2i，右儿子在2i+1，父节点在i/2
//		3. 基于数组存储的堆，堆的大小需要事先估算
// 堆序性: 任意节点应该小于或等于它的后裔节点，快速从根上找到最小元(小顶堆)，堆有关序的信息很少，查找某个元素，需要对堆进行线性查找！
// 操作:
//		1. Insert: 在堆中寻求合适位置，插入元素X。依次上滤待插入X元素，不破坏堆序性寻找插入位置，直至到根节点
//		2. DelMin:
//			a) 删除最小元素，数组1号位置存储的即最小元素;
//			b) 删除后，有空洞，需要从删除元素左右子节点挑选一个填补空洞，依次迭代下滤空洞;
//			c) 可能存在下滤到非最后一个节点，故应先将最后节点填补空洞，再下滤，直至找到一个满足堆序性的位置；
// 相关应用和算法：
// 		1. 贪婪算法: 反复求最小/最大元素
//		2. 选择问题: 找出N个元素列表中的第k大元素(构建一个k个节点的小顶堆，每次和堆顶元素相比，小跳过，大则淘汰堆顶元素，填补新入元素)
//		3. 事件模拟/任务调度
//		4. 堆排序
// 其他补充：
//		1. 满二叉树(perfect binary tree): 一棵高位h的二叉树，满足拥有2^(h+1)-1个节点，称之为满二叉树(每个叉都是满的)
//		2. 满二叉树树是完全二叉树，但反之不一定
//		3. d-堆: 所有节点都有d个儿子，相比二叉堆，其深度要浅得多，但需要从d中找到最小者，需要d-1次比较(类似b-树的作用)
//		4. 堆的合并，比较困难，由此产生了三类ADT: 左式堆、斜堆、二项队列
//		5. 左式堆: 合并容易(左儿子至少与右儿子一样长）
type PriorityQueue interface {
	Insert(x int)
	Delete() int
	Print()
}

//MinHeap 小顶堆
type MinHeap struct {
	Cap   int
	Len   int
	Elems []int
}

//NewMinHeap 堆初始化
func NewHeap(cap int) *MinHeap {
	h := &MinHeap{
		Cap:   cap,
		Elems: make([]int, cap+1),
	}
	h.Elems[0] = -1
	return h
}

//Insert 向堆中插入元素n
func (h *MinHeap) Insert(x int) {
	// 1. 边界条件
	if x <= 0 {
		panic("x must be positive integer")
	}
	if h.Len >= h.Cap {
		panic("MinHeap is full or x")
	}

	// 2. 空穴上滤: 比较堆E[h.len+1]位置和其父节点元素大小，依次迭代到根位置
	p := h.Len + 1
	h.Elems[p], h.Len = x, h.Len+1
	for ; h.Elems[p] < h.Elems[p/2]; p = p / 2 {
		h.Elems[p/2], h.Elems[p] = h.Elems[p], h.Elems[p/2]
	}
}

//Print 层序遍历, 打印堆中相关元素
func (h *MinHeap) Print() {
	for i := 1; i <= h.Len; i++ {
		fmt.Printf("%d ", h.Elems[i])
	}
	fmt.Printf("\n")
}

//Delete 从堆中移除堆顶元素
func (h *MinHeap) Delete() (elem int) {
	// 1. 边界条件
	if h.Len == 0 {
		panic("Heap is empty")
	}

	// 2. 下滤，移除E[1]元素, 利用E[h.len+1]填补，同时下滤根节点元素
	pop := h.Elems[1]
	h.Elems[1], h.Len = h.Elems[h.Len], h.Len-1
	for i, child := 1, 0; 2*i < h.Len; i = child {
		// 判断左右元素的小者
		switch {
		case h.Elems[2*i] < h.Elems[2*i+1]:
			child = 2 * i
		default:
			child = 2*i + 1
		}
		// 判断当前i节点元素是否满足插入条件, 或者选择对应的child子堆进行下滤
		if h.Elems[i] > h.Elems[child] {
			h.Elems[i], h.Elems[child] = h.Elems[child], h.Elems[i]
		} else {
			break
		}
	}

	return pop
}
