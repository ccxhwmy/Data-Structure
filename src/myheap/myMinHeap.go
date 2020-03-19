package myheap

import "errors"

type MinHeap struct{
	data []int
	count int
	capacity int
}


func NewMinHeap( capacity int ) *MinHeap {
	return &MinHeap{
		data:  make( []int, capacity + 1 ),
		count: 0,
		capacity: capacity,
	}
}

func NewMinHeapArr( arr []int, n int ) *MinHeap {
	retHeap := &MinHeap{
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

func ( mh *MinHeap )Size() int {
	return mh.count
}

func ( mh *MinHeap )IsEmpty() bool {
	return mh.count == 0
}

func ( mh *MinHeap )__shiftUp( k int ) {
	e := mh.data[k]
	for k > 1 && mh.data[k/2] > e {
		mh.data[k] = mh.data[k/2]
		k /= 2
	}
	mh.data[k] = e
}

func ( mh *MinHeap )Insert( n int ) {
	if mh.count + 1 > mh.capacity {
		panic( "heap is full" )
	}
	mh.data[mh.count+1] = n
	mh.count++
	mh.__shiftUp( mh.count )
}

func ( mh *MinHeap )__shiftDown( k int ) {
	e := mh.data[k]
	for 2*k <= mh.count {
		j := 2*k
		if j+1 <= mh.count && mh.data[j+1] < mh.data[j] {
			j++
		}
		if e <= mh.data[j] {
			break
		}
		mh.data[k] = mh.data[j]
		k = j
	}
	mh.data[k] = e
}

func ( mh *MinHeap )ExtractMax() ( int, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	ret := mh.data[1]
	mh.data[1], mh.data[mh.count] = mh.data[mh.count], mh.data[1]
	mh.count--
	mh.__shiftDown( 1 )
	return ret, nil
}

func ( mh *MinHeap )GetMin() ( int, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	return mh.data[1], nil
}

