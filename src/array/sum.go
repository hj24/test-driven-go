package array

/*
数组有的大小也属于类型的一部分，
如果你尝试将 [4]int 作为 [5]int 类型的参数传入函数，是不能通过编译的。
因为这个原因，所以数组比较笨重，大多数情况下我们都不会使用它。
*/
func Sum(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SliceSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(numtosum ...[]int) (sums []int) {
	lenofargs := len(numtosum)
	sums = make([]int, lenofargs)
	for i, v := range numtosum {
		sums[i] = SliceSum(v)
	}
	return sums
}

func SumAllTail(arrs ...[]int) (sums []int) {
	for _, v := range arrs {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SliceSum(v[1:]))
		}
	}
	return sums
}
