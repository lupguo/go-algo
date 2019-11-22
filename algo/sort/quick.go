package sort

func Quick(A []int, n int) (sortedList []int) {
	// 1. 元素是否仅单个值
	if len(A) < 2 {
		return A
	}

	// 2. 选择合适的Pivot值，交换至最后一个值
	Pivot(A, n)

	// 3. 基于选取的pivot值进行分区，确保A[:p], A[p], A[p+1:]三个集合顺序关系
	p := Partition(A, len(A))

	// 4. 分别针对左右分区进行快排
	Quick(A[:p], len(A[:p]))
	Quick(A[p+1:], len(A[p+1:]))

	return A
}

//Partition 将列表A，基于pivot进行分区
// 注意:
// 		第一个for, l < p: l游标持续的往右推进，直至l游到最后，或者遇到一个比pivot点大的元素
// 			A[l] <= A[p]: 等号是为了在列表元素都相同情况下，左游标会继续前进，不至于循环无法退出，同时确保排序稳定性
// 		第二个for, r > 0: r游标持续往左推进，直至r到列表左边界，或者遇到一个比pivot点小的元素
// 			A[p] <= A[r], 同理，若右值和分区点值相等的情况，右游标继续左移，同时这样保证了排序的稳定性
//		退出条件: l >= r: l和r游标已经针对所有待分区区间都分割好了, l即分区点位置
func Partition(A []int, n int) (pivot int) {
	// 设定游标 l, r，分区点下标p
	l, r, p := 0, n-2, n-1

	// l右移, r左移, 当l,r相遇, 循环退出
	for {
		for l < p && A[l] < A[p] {
			l++
		}
		for r > 0 && A[r] >= A[p] {
			r--
		}
		if l >= r {
			break
		}
		// 交换l游标和r游标元素
		A[l], A[r] = A[r], A[l]
	}
	A[l], A[p] = A[p], A[l]
	return l
}

//Pivot 中值法确定pivot分区点，挑选合适的pivot，将其值和列表最后一个元素交换
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
