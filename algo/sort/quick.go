package sort

//QSortVer1 快排
//	1. 选取一个参考点Pivot，利用分区点将数组分成两个区间
//	2. 分别递归的对区分的两个区间进行快排
func QSortVer1(A []int, n int) (sortedList []int) {
	// 1. 递归退出条件，元素是否仅单个值
	if len(A) < 2 {
		return A
	}
	// 2. 选择合适的Pivot值，同时交换至最后一个值
	Pivot(A, n)
	// 3. 基于选取的pivot值找到确定的分区点p, 利用分区点进行分区，确保A[:p], A[p], A[p+1:]三个集合顺序关系
	p := PartitionVer1(A, len(A))
	// 4. 分别针对左右分区进行快排
	QSortVer1(A[:p], len(A[:p]))
	QSortVer1(A[p+1:], len(A[p+1:]))
	return A
}

//PartitionVer1 分区方法1，游标探测法 - (思想类似于选择排序，分成小于pivot左区间，以及探测区间)
// 	1. 循环从最i元素迭代开始，往前进行j游标探测，在探测区间内，寻找小于pivot值的元素；
// 		循环内若j往前探测区域内，探测到一个比pivot小的值，则进行A[i]与A[j]交换，同时i往前一个元素；
//	2. 迭代完整个A列表后，i所在位置左侧都比pivot小，右侧包括自己则不小于pivot，因此交换A[i]与A[p]的值，并返回找到的pivot下标
//	参考: https://time.geekbang.org/column/article/41913
func PartitionVer1(A []int, n int) (index int) {
	i, p := 0, n-1
	for j := i; j < p; j++ {
		if A[j] < A[p] {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i], A[p] = A[p], A[i]
	return i
}

//Pivot 【中值法】确定pivot分区点，挑选合适的pivot，将其值和列表最后一个元素交换
func Pivot(A []int, n int) {
	var pivot int
	min, mid, max := 0, n/2, n-1
	if A[min] < A[mid] {
		switch {
		case A[mid] < A[max]:
			pivot = mid
		case A[min] < A[max]:
			pivot = max
		default:
			pivot = min
		}
	} else {
		// A[mid] <= A[min]
		switch {
		case A[max] < A[mid]:
			pivot = mid
		case A[min] < A[max]:
			pivot = min
		default:
			pivot = max
		}
	}
	//return len(A) - 1
	A[pivot], A[n-1] = A[n-1], A[pivot]
}

// v2
func QSortVer2(A []int, n int) (sortedList []int) {
	if n < 2 {
		return A
	}
	p := PartitionVer2(A, n)
	QSortVer2(A[:p], len(A[:p]))
	QSortVer2(A[p+1:], len(A[p+1:]))
	return A
}

// PartitionVer2 分区方法2，左右游标l,r居中夹逼法
//		即针对列表A，基于pivot(最后一个元素)进行分区，同时返回指定分区值的下标index
//
// 时间复杂度：T(N)=N + 2T(N/2)，N为Partition的时间开销，迭代可求得时间复杂度为O(N*log(N))
// 空间复杂度：Partition期间，仅O(1)空间开销，属于原地排序
// 不是稳定的排序算法，举例[6,5,4,6,2,5]，当第一次A[l]与A[r]交换后，两个6的前后位置已经发生了变化
//
// 	Tips: 选择分区方法1，容易理解一些，代码也较为精简
func PartitionVer2(A []int, n int) int {
	l, r, p := 0, n-2, n-1
	pivot := A[p]
	for {
		// l->
		for A[l] <= pivot && l < p {
			l++
		}
		// l == p
		if l == p {
			break
		}
		// <-r
		for A[r] >= pivot && r > l {
			r--
		}
		// l == r
		if l == r {
			break
		}
		// swap l, r
		A[l], A[r] = A[r], A[l]
	}
	A[l], A[p] = A[p], A[l]
	return l
}
