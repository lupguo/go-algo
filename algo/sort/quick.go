package sort

//QSort 快排
func QSort(A []int, n int) (sortedList []int) {
	// 1. 递归退出条件，元素是否仅单个值
	if len(A) < 2 {
		return A
	}

	// 2. 选择合适的Pivot值，同时交换至最后一个值
	Pivot(A, n)

	// 3. 基于选取的pivot值进行分区，确保A[:p], A[p], A[p+1:]三个集合顺序关系
	//p := Partition(A, len(A))
	p := PartitionX(A, len(A))

	// 4. 分别针对左右分区进行快排
	QSort(A[:p], len(A[:p]))
	QSort(A[p+1:], len(A[p+1:]))

	return A
}

//Partition 分区方法一，左右游标l,r居中夹逼法，即针对列表A，基于pivot(最后一个元素)进行分区，同时返回指定分区值的下标index
// 注意:
// 		第一个for, l < p: l游标持续的往右推进，直至l游到最后，或者遇到一个比pivot点大的元素
// 			A[l] <= A[p]: 等号是为了在列表元素都相同情况下，左游标会继续前进，不至于循环无法退出，同时确保排序稳定性
// 		第二个for, r > 0: r游标持续往左推进，直至r到列表左边界，或者遇到一个比pivot点小的元素
// 			A[p] <= A[r], 同理，若右值和分区点值相等的情况，右游标继续左移，同时这样保证了排序的稳定性
//		退出条件: l >= r: l和r游标已经针对所有待分区区间都分割好了, l即分区点位置
// 时间复杂度：T(N)=N + 2T(N/2)，N为Partition的时间开销，迭代可求得时间复杂度为O(N*log(N))
// 空间复杂度：Partition期间，仅O(1)空间开销，属于原地排序
// 不是稳定的排序算法，举例[6,5,4,6,2,5]，当第一次A[l]与A[r]交换后，两个6的前后位置已经发生了变化
// Tips: 选择分区方法2，容易理解一些，代码也较为精简
func Partition(A []int, n int) (index int) {
	// 设定游标 l, r，分区点下标p(哨兵)
	l, r, p := 0, n-2, n-1

	// l右移, r左移, 当l>=r, 循环退出
	for {
		for l < p && A[l] < A[p] {
			l++
		}
		for r > 0 && A[r] >= A[p] { // Tips: 此处的A[r]>=A[p]，不能替换为A[r]>A[p]，因为当A[l]=A[r]的情况，会出现死循环
			r--
		}
		if l >= r { // Tips: 此处的>=号，不能替换为>号，当存在[5,2]的情况，因为当l=r=0的情况，会出现死循环
			break
		}
		// 交换l游标和r游标元素
		A[l], A[r] = A[r], A[l]
	}
	// 最后l所在的元素位置，即为pivot的位置，将pivot兑换到该位置
	A[l], A[p] = A[p], A[l]
	return l
}

//PartitionX 分区方法二，快游标探测法；
// 	1. 循环从最i元素迭代开始，往前进行j游标探测，在探测区间内，寻找小于pivot值的元素；
// 		循环内若j往前探测区域内，探测到一个比pivot小的值，则进行A[i]与A[j]交换，同时i往前一个元素；
//	2. 迭代完整个A列表后，i所在位置左侧都比pivot小，右侧包括自己则不小于pivot，因此交换A[i]与A[p]的值，并返回找到的pivot下标
//	参考: https://time.geekbang.org/column/article/41913
func PartitionX(A []int, n int) (index int) {
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
