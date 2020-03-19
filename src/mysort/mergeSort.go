package mysort

import "mypublic"

func __merge( arr []int, l, mid, r int )  {
	aux := make( []int, r-l+1 )
	i, j := l, mid + 1
	for k := 0; k <= r-l; k++ {
		if i > mid {
			aux[k] = arr[j]
			j++
		} else if j > r {
			aux[k] = arr[i]
			i++
		} else if arr[i] < arr[j] {
			aux[k] = arr[i]
			i++
		} else {
			aux[k] = arr[j]
			j++
		}
	}
	copy( arr[l:r+1], aux )
}

/* 递归使用归并排序,对arr[l...r]的范围进行排序 */
func __mergeSort( arr []int, l, r int ) {
	if l >= r {
		return
	}
	/* 减小溢出的可能 */
	mid := l + (r - l) / 2
	__mergeSort( arr, l, mid )
	__mergeSort( arr, mid+1, r )
	__merge( arr, l, mid, r )
}

func __mergeSort2( arr []int, l, r int ) {
	if r - l <= 15 {
		insertSortRange( arr, l, r )
		return
	}

	/* 减小溢出的可能 */
	mid := l + (r - l) / 2
	__mergeSort2( arr, l, mid )
	__mergeSort2( arr, mid+1, r )

	/* 对于arr[mid] <= arr[mid+1]的情况,不进行merge, 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失 */
	if arr[mid] > arr[mid+1] {
		__merge( arr, l, mid, r )
	}
}

func mergeSort( arr []int, n int ) {
	__mergeSort2( arr, 0, n-1 )
}


// Merge Sort BU 也是一个O(nlogn)复杂度的算法，虽然只使用两重for循环
// 所以，Merge Sort BU也可以在1秒之内轻松处理100万数量级的数据
// 不要轻易根据循环层数来判断算法的复杂度，Merge Sort BU就是一个反例
func mergeSortBU( arr []int, n int ) {
	/* 对于小数组, 使用插入排序优化 */
	for i := 0; i < n; i += 16 {
		insertSortRange( arr, i, mypublic.MyMin( i + 15, n - 1 ) )
	}
	for sz := 16; sz < n; sz += sz {
		for i := 0; i < n - sz; i += sz + sz {
			if arr[i+sz-1] > arr[i+sz] {
				__merge( arr, i, i+sz-1, mypublic.MyMin( i+sz+sz-1, n-1 ) )
			}
		}
	}
}





/////////////////////////////////////////////////////////////////////////




func mergeHelper( arr []int, l, mid, r int )  {
	aux := make( []int, ( r-l+1 ) )
	i, j := l, mid + 1
	k := 0
	for ; k < r-l+1; k++ {
		if i > mid {
			aux[k] = arr[j]
			j++
		} else if j > r {
			aux[k] = arr[i]
			i++
		} else if arr[i] < arr[j] {
			aux[k] = arr[i]
			i++
		} else {
			aux[k] = arr[j]
			j++
		}
	}
	copy( arr[l:r+1], aux )
}

func mergeSortHelper( arr []int, l, r int ) {
	if r - l <= 15 {
		insertSortRange( arr, l, r )
		return
	}
	mid := l + ( r - l ) / 2
	mergeSortHelper( arr, l, mid )
	mergeSortHelper( arr, mid+1, r )
	if arr[mid] > arr[mid+1] {
		mergeHelper( arr, l, mid, r )
	}
}

func MergeSort( arr []int, n int ) {
	mergeSortHelper( arr, 0, n-1 )
}

