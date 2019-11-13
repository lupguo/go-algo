package sort

//Select 选择排序, 每第i趟迭代从未排好序的A[i+1,n-1]列表中，挑选最小值，和当前第i个元素交换
// 原地排序,  最好,最坏都是O(N^2)
// 不是稳定的排序算法：因为针对5,8,5,2,9，第一个5会与2交换，排序前后相同元素的位置变化了
func Select(A []int, n int) (sortedList []int)  {
	for i:=0; i<n; i++ {
		min := i
		for j:=i+1; j<n; j++ {
			if A[min] > A[j] { // 从未排好序列表中，找到更小值，则交换
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
	return A
}
