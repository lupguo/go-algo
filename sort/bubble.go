package sort

//Bubble 冒泡排序：原地排序、稳定、O(n^2)
func Bubble(A []int, n int) (sortedList []int) {
	for i := 0; i < n-1; i++ {
		swap := false
		// 内层循环，每趟进行相邻元素的比较，依次冒泡出最大元素(全部利用j游标进行相邻元素交换)
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
				swap = true
			}
		}
		// 若内层遍历过程中，无交换则表示列表已排序完成，退出外层循环
		if !swap {
			break
		}
	}
	return A
}
