package myheap

import "errors"

type MaxHeap struct{
	data []int
	count int
	capacity int
}

/* 创建一个空堆, 可容纳capacity个元素 */
func NewMaxHeap( capacity int ) *MaxHeap {
	return &MaxHeap{
		data:  make( []int, capacity + 1 ),
		count: 0,
		capacity: capacity,
	}
}

/* 通过一个给定数组创建一个最大堆, 该构造堆的过程, 时间复杂度为O(n) */
func NewMaxHeapArr( arr []int, n int ) *MaxHeap {
	retHeap := &MaxHeap{
		data:     make( []int, n + 1 ),
		count:    0,
		capacity: n,
	}
	for i := 0; i < n; i++ {
		retHeap.data[i+1] = arr[i]
	}
	retHeap.count = n
	for i := retHeap.count / 2; i >= 1; i-- {
		retHeap.__shiftDown( i )
	}
	return retHeap
}

/* 返回堆中的元素个数 */
func ( mh *MaxHeap )Size() int {
	return mh.count
}

/* 返回一个布尔值, 表示堆中是否为空 */
func ( mh *MaxHeap )IsEmpty() bool {
	return mh.count == 0
}

func ( mh *MaxHeap )__shiftUp( k int ) {
	e := mh.data[k]
	for k > 1 && mh.data[k/2] < e {
		mh.data[k] = mh.data[k/2]
		k /= 2
	}
	mh.data[k] = e
}

/* 向最大堆中插入一个新的元素 item */
func ( mh *MaxHeap )Insert( n int ) {
	if mh.count + 1 > mh.capacity {
		panic( "heap is full" )
	}
	mh.data[mh.count+1] = n
	mh.count++
	mh.__shiftUp( mh.count )
}

func ( mh *MaxHeap )__shiftDown( k int ) {
	e := mh.data[k]
	for 2*k <= mh.count {
		j := 2*k
		if j+1 <= mh.count && mh.data[j+1] > mh.data[j] {
			j++
		}
		if e >= mh.data[j] {
			break
		}
		mh.data[k] = mh.data[j]
		k = j
	}
	mh.data[k] = e
}

/* 从最大堆中取出堆顶元素, 即堆中所存储的最大数据 */
func ( mh *MaxHeap )ExtractMax() ( int, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	ret := mh.data[1]
	mh.data[1], mh.data[mh.count] = mh.data[mh.count], mh.data[1]
	mh.count--
	mh.__shiftDown( 1 )
	return ret, nil
}

/* 获取最大堆中的堆顶元素 */
func ( mh *MaxHeap )GetMax() ( int, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	return mh.data[1], nil
}
