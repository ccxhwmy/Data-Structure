package mysort

/* 插入排序 */
func insertSort( arr []int, n int ) {
	for i := 1; i < n; i++ {
		e := arr[i]
		var j int
		for j = i; j > 0 && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

/* 对某一段范围进行插入排序 */
func insertSortRange( arr []int, l, r int )  {
	for i := l + 1; i <= r; i++ {
		e := arr[i]
		var j int
		for j = i; j > l && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

//////////////////////////////////////////////////////////////////////////////


func InsertSort( arr []int, n int ) {
	for i := 1; i < n; i++ {
		e := arr[i]
		j := i
		for ; j > 0 && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}


