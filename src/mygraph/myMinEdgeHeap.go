package mygraph

import "errors"

type MinEdgeHeap struct{
	data []*Edge
	count int
	capacity int
}


func NewMinEdgeHeap( capacity int ) *MinEdgeHeap {
	return &MinEdgeHeap{
		data:  make( []*Edge, capacity + 1 ),
		count: 0,
		capacity: capacity,
	}
}

func ( mh *MinEdgeHeap )Size() int {
	return mh.count
}

func ( mh *MinEdgeHeap )IsEmpty() bool {
	return mh.count == 0
}

func ( mh *MinEdgeHeap )shiftUp( k int ) {
	e := mh.data[k]
	for k > 1 && mh.data[k/2].weight > e.weight {
		mh.data[k] = mh.data[k/2]
		k /= 2
	}
	mh.data[k] = e
}

func ( mh *MinEdgeHeap )Insert( n int ) {
	if mh.count + 1 > mh.capacity {
		panic( "heap is full" )
	}
	mh.data[mh.count+1] = n
	mh.count++
	mh.shiftUp( mh.count )
}

func ( mh *MinEdgeHeap )shiftDown( k int ) {
	e := mh.data[k]
	for 2*k <= mh.count {
		j := 2*k
		if j+1 <= mh.count && mh.data[j+1].weight < mh.data[j].weight {
			j++
		}
		if e.weight <= mh.data[j].weight {
			break
		}
		mh.data[k] = mh.data[j]
		k = j
	}
	mh.data[k] = e
}

func ( mh *MinEdgeHeap )ExtractMax() ( *Edge, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	ret := mh.data[1]
	mh.data[1], mh.data[mh.count] = mh.data[mh.count], mh.data[1]
	mh.count--
	mh.shiftDown( 1 )
	return ret, nil
}

func ( mh *MinEdgeHeap )GetMin() ( *Edge, error ) {
	if mh.count <= 0 {
		return 0, errors.New( "heap is nil" )
	}
	return mh.data[1], nil
}

