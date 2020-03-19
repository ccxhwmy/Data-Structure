package mysort

import (
	"math/rand"
	"time"
)

/* 对arr[l...r]部分进行partition操作 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p] */
func __partition( arr []int, l, r int ) int {
	/* 防止数组有序性过强，随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot */
	swapIndex := rand.Intn( r-l ) + l
	arr[l], arr[swapIndex] = arr[swapIndex], arr[l]

	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}


/* 双路快速排序的partition */
func __partition2( arr []int, l, r int ) int {
	/* 防止数组有序性过强，随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot */
	swapIndex := rand.Intn( r-l ) + l
	arr[l], arr[swapIndex] = arr[swapIndex], arr[l]
	v := arr[l]
	i, j := l + 1, r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l + 1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}


/* 对arr[l...r]部分进行快速排序 */
func __quickSort( arr []int, l, r int ) {
	/* 对于小规模数组, 使用插入排序进行优化 */
	if r - l <= 15 {
		insertSortRange( arr, l, r )
		return
	}
	p := __partition( arr, l, r )
	__quickSort( arr, l, p - 1 )
	__quickSort( arr, p + 1, r )
}

func quickSort( arr []int, n int ) {
	rand.Seed( time.Now().Unix() )
	__quickSort( arr, 0, n-1 )
}


func __quickSort3Ways( arr []int, l, r int ) {
	if r - l <= 15 {
		insertSortRange( arr, l, r )
		return
	}

	/* 防止数组有序性过强，随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot */
	swapIndex := rand.Intn( r-l ) + l
	arr[l], arr[swapIndex] = arr[swapIndex], arr[l]
	v := arr[l]
	/* arr[l+1...lt] < v, arr[lt+1...i) == v, arr[gt...r] > v  */
	lt, gt, i := l, r + 1, l + 1
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	__quickSort3Ways( arr, l, lt-1 )
	__quickSort3Ways( arr, gt, r )
}

func quickSort3Ways( arr []int, n int ) {
	rand.Seed( time.Now().Unix() )
	__quickSort3Ways( arr, 0, n-1 )
}


////////////////////////////////////////////////////////////////////////////////////////////////


func quickSortthreeWays( arr []int, l, r int )  {
	if r - l <= 15 {
		insertSortRange( arr, l, r )
		return
	}
	randIndex := rand.Intn( r-l+1 ) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]
	v := arr[l]
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < v {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > v {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	quickSortthreeWays( arr, l, lt-1 )
	quickSortthreeWays( arr, gt, r )
}



func QuickSort3Ways( arr []int, n int )  {
	rand.Seed( time.Now().Unix() )
	quickSortthreeWays( arr, 0, n-1 )
}




