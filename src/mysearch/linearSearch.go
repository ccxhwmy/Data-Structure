package mysearch


// 线性查找法, 实现lower_bound
// 即在一个有序数组arr中, 寻找大于等于target的元素的第一个索引
// 如果存在, 则返回相应的索引index
// 否则, 返回arr的元素个数 n
func LowerBound( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0" )
	}

	for i := 0; i < n; i++ {
		if arr[i] >= target {
			return i
		}
	}
	return n
}


// 线性查找法, 实现upper_bound
// 即在一个有序数组arr中, 寻找大于target的元素的第一个索引
// 如果存在, 则返回相应的索引index
// 否则, 返回arr的元素个数 n
func UpperBound( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0" )
	}

	for i := 0; i < n; i++ {
		if arr[i] > target {
			return i
		}
	}
	return n
}
