package sort

//Insert 插入排序，分成未排序和已排序区间，原地排序、稳定排序(相同元素可以插其后)、O(N^2)
//	Tips: 插入排序需要挪动从[待插入, 已排序]区间内的元素数据
func Insert(A []int, n int) (sortedList []int) {
	for i := 1; i < n; i++ {
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
