package sort

//MSort 归并 - 分治: 先分区成两个列表，然后Merge有序列表
func MSort(A []int, n int) (sortedList []int) {
	// 1. 递归退出条件
	if n < 2 {
		return A
	}

	// 2. 分区成两个列表，然后分别对两个列表进行归并排序
	lSort := MSort(A[:n/2], n/2)
	rSort := MSort(A[n/2:], n-n/2)

	// 3. 治的过程
	return Merge(lSort, rSort)
}

//Merge 治的过程,合并两个有序列表
func Merge(A []int, B []int) (sortedList []int) {
	lenA, lenB := len(A), len(B)

	// 1. 申请一个M=lenA+lb的列表, pa, pb游标，分别指向A, B首地址
	M := make([]int, lenA+lenB)
	pa, pb := 0, 0

	// 2. 依次比较A[pa]和B[pb]值，基于pa, pb游标，任意一个完成，退出循环
	for pa < lenA && pb < lenB {
		if A[pa] < B[pb] {
			M[pa+pb] = A[pa]
			pa++
		} else {
			M[pa+pb] = B[pb]
			pb++
		}
	}

	// 3. A, B列表中没有迭代完成的，直接将剩余(A[pa:] or B[pb:])的内容copy到M
	if pa > lenA-1 { // pa迭代完, 补B剩余
		copy(M[pa+pb:], B[pb:])
	} else {
		copy(M[pa+pb:], A[pa:])
	}
	return M
}
