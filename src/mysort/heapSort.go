package mysort

import (
	"fmt"
	"myheap"
)

func heapSort1( arr []int, n int ) {
	var err error
	maxheap := myheap.NewMaxHeap( n )
	for i := 0; i < n; i++ {
		maxheap.Insert( arr[i] )
	}

	for i := n-1; i >= 0; i-- {
		arr[i], err = maxheap.ExtractMax()
		if err != nil {
			fmt.Println( err )
		}
	}
}

func heapSort2( arr []int, n int ) {
	var err error
	maxheap := myheap.NewMaxHeapArr( arr, n )
	for i := n-1; i >= 0; i-- {
		arr[i], err = maxheap.ExtractMax()
		if err != nil {
			fmt.Println( err )
		}
	}
}

func heapSortUsingMinHeap( arr []int, n int )  {
	var err error
	minheap := myheap.NewMinHeap( n )
	for i := 0; i < n; i++ {
		minheap.Insert( arr[i] )
	}

	for i := 0; i < n; i++ {
		arr[i], err = minheap.ExtractMax()
		if err != nil {
			fmt.Println( err )
		}
	}
}

func heapSortUsingMinHeap2( arr []int, n int )  {
	var err error
	minheap := myheap.NewMinHeapArr( arr, n )
	for i := 0; i < n; i++ {
		arr[i], err = minheap.ExtractMax()
		if err != nil {
			fmt.Println( err )
		}
	}
}

func __shiftDown( arr []int, n, k int ) {
	e := arr[k]
	for 2*k+1 < n {
		j := 2*k+1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if arr[j] <= e {
			break
		}
		arr[k] = arr[j]
		k = j
	}
	arr[k] = e
}

/* 不使用一个额外的最大堆, 直接在原数组上进行原地的堆排序 */
func heapSort3( arr []int, n int ) {

	// 此时我们的堆是从0开始索引的
	// 从(最后一个元素的索引-1)/2开始
	// 最后一个元素的索引 = n-1
	for i := ( n-1-1 ) / 2; i >= 0; i-- {
		__shiftDown( arr, n, i )
	}
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		__shiftDown( arr, i, 0 )
	}
}

































///////////////////////////////////////////////////

func shiftDown( arr []int, n, k int )  {
	e := arr[k]
	for 2*k+1 < n {
		j := 2*k+1
		if j + 1 < n && arr[j] < arr[j+1] {
			j = j + 1
		}
		if e >= arr[j] {
			break
		}
		arr[k] = arr[j]
		k = j
	}
	arr[k] = e
}

func HeapSort( arr []int, n int ) {
	for i := (n-2)/2; i >= 0; i-- {
		shiftDown( arr, n, i )
	}

	for i := n-1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown( arr, i, 0 )
	}
}
