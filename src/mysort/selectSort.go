package mysort

/* 选择排序 */
func selectSort( arr []int, n int ) {
	for i := 0; i < n; i++ {
		// 寻找[i, n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

/* 遍历时分别对最大值与最小值进行选择排序 */
func selectSort2( arr []int, n int ) {
	l, r := 0, n-1
	for l < r {
		minIndex := l
		maxIndex := r
		if arr[minIndex] > arr[maxIndex] {
			arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
		}
		for i := l + 1; i < r; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			} else if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		arr[minIndex], arr[l] = arr[l], arr[minIndex]
		arr[maxIndex], arr[r] = arr[r], arr[maxIndex]
		l++
		r--
	}
}


/////////////////////////////////////////////////////////////////////////////////////////////////

func SelectSort( arr []int, n int ) {
	l, r := 0, n-1
	for l < r {
		minIndex := l
		maxIndex := r
		if arr[minIndex] > arr[maxIndex] {
			arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
		}
		for i := l; i < r; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			} else if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		arr[l], arr[minIndex] = arr[minIndex], arr[l]
		arr[r], arr[maxIndex] = arr[maxIndex], arr[r]
		l++
		r--
	}
}


