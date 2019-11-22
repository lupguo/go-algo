package sort

func Insert(A []int) (sortedList []int) {
	for i := 1; i < len(A); i++ {
		// j初始化指向i位置, tmp记录当前轮待插入元素值
		j, tmp := i, A[i]
		// 迭代前p个有序元素，依次将A[j-1]和待插入元素值比较
		// 当比待交换元素大时，由于是数组，需要依次后移A[j-1]的值，腾挪空间给待插入元素
		for ; j > 0 && A[j-1] > tmp; j-- {
			A[j] = A[j-1]
		}
		// j即待插入的位置，插入tmp值
		A[j] = tmp
	}
	return A
}
